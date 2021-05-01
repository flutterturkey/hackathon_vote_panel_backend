package controllers

import (
	"github.com/labstack/echo/v4"
	"goBoilterplate/app/helpers"
	"goBoilterplate/app/models"
	"goBoilterplate/config"
	"net/http"
	"time"
)

// Index godoc
// @Summary Home Page
// @Description Display Home Page
// @Tags Home
// @Produce  json
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router / [get]
func Index(c echo.Context) error {
	response := models.ProjectsResponse{Error: "Not Found", Data: nil}
	return c.JSON(http.StatusNotFound, response)
}

// Login godoc
// @Summary Login
// @Description Login User in API
// @Tags Auth
// @Produce  json
// @Param email formData string true "Email"
// @Param password formData string true "Password"
// @Success 200 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /api/login [post]
func Login(c echo.Context) error {
	login := models.Login{}
	login.Email = c.FormValue("email")
	login.Password = c.FormValue("password")

	err := helpers.Validate(&login)
	if err != nil {
		return c.JSON(422, err)
	}

	user := models.AuthLogin(login.Email, login.Password)
	if user != nil {
		token, err := helpers.AuthMakeToken(user)
		if err != nil {
			return c.JSON(500, "Server Error")
		}

		err = config.DB.Model(&user).Where("email = ? and password = ?", login.Email, login.Password).Update("login", time.Now()).Error

		if err != nil {
			return c.JSON(http.StatusOK, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
		}
		response := models.BaseResponse{Data: map[string]string{"token": token}}
		return c.JSON(http.StatusOK, response)
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Error: true, Data: models.Message{Message: "E-mail veya şifre yanlış!"}})
}

// Logout godoc
// @Summary Logout
// @Description User Logout
// @Tags Auth
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {string} string
// @Failure 401 {string} string
// @Router /api/logout [get]
func Logout(c echo.Context) error {
	user := helpers.AuthGetUser(c)
	if user == nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Error: true, Data: models.Message{Message: "Hacı yanlış bilgilerle deneme işte"}})
	}

	err := config.DB.Model(&user).Where("id = ?", user.ID).Update("logout", time.Now()).Error

	if err != nil {
		return c.JSON(http.StatusOK, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{Data: models.Message{Message: "Çıkış başarılı!"}})

}

// Ready godoc
// @Summary Ready status
// @Description Is api ready
// @Tags Home
// @Produce html
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /ready [get]
func Ready(c echo.Context) error {
	return c.HTML(200, "OK")
}
