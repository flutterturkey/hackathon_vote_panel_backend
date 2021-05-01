package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/lib/pq"
	"goBoilterplate/app/helpers"
	"goBoilterplate/app/models"
	"goBoilterplate/config"
	"net/http"
	"strconv"
)

// Index godoc
// @Summary Get Projects
// @Description Display Projects
// @Tags Projects
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /projects [get]
func ProjectList(c echo.Context) error {
	tmpUser := helpers.AuthGetUser(c)
	if tmpUser == nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Error: true, Data: models.Message{Message: "Oturum açman lazım cano"}})

	}

	var projects []models.Projects
	var user models.User

	err := config.DB.Find(&projects).Error
	if err != nil {
		return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
	}

	err = config.DB.Where("id = ?", tmpUser.ID).First(&user).Error

	if err != nil {
		return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
	}

	for i := 0; i < len(projects); i++ {
		if itemInArray(int(user.ID), projects[i].Voters) {
			projects[i].IsLiked = true
		}
	}

	return c.JSON(http.StatusOK, models.ProjectsResponse{Data: projects})

}

// Index godoc
// @Summary Project Detail
// @Description Display Project Detail
// @Tags Projects
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /projects/{id}/ [get]
// @ID get-string-by-int
// @Param id path int true "Project"
func ProjectDetail(c echo.Context) error {
	tmpUser := helpers.AuthGetUser(c)
	if tmpUser == nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Error: true, Data: models.Message{Message: "Oturum açman lazım cano"}})
	}
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
	}

	var project models.Project
	var user models.User

	err = config.DB.Where("id = ?", id).First(&project).Error
	if err != nil {
		return c.JSON(http.StatusOK, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
	}

	err = config.DB.Where("id = ?", tmpUser.ID).First(&user).Error

	if err != nil {
		return c.JSON(404, err)
	}

	if itemInArray(int(user.ID), project.Voters) {
		project.IsLiked = true
	}

	return c.JSON(http.StatusOK, models.ProjectDetailResponse{Data: project})

}

// Index godoc
// @Summary Project Upvote
// @Description Project Upvote
// @Tags Projects
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /projects/{id}/ [post]
// @ID get-string-by-int
// @Param id path int true "Project"
func ProjectUpvote(c echo.Context) error {
	tmpUser := helpers.AuthGetUser(c)
	if tmpUser == nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Error: true, Data: models.Message{Message: "Oturum açman lazım cano"}})
	}

	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(404, err)
	}

	var project models.Project
	var user models.User

	err = config.DB.Where("id = ?", id).First(&project).Error

	if err != nil {
		return c.JSON(404, err)
	}

	err = config.DB.Where("id = ?", tmpUser.ID).First(&user).Error

	if err != nil {
		return c.JSON(404, err)
	}

	if len(user.Votes) >= models.UserVoteLimit {
		return c.JSON(http.StatusBadRequest,
			models.BaseResponse{Error: true, Data: models.Message{Message: "Oylama limitiniz doldu! Oylarınızı geri alarak tekrar deneyebilirsiniz."}})
	}

	if user.TeamName == project.TeamName {
		return c.JSON(http.StatusTeapot, models.BaseResponse{Error: true, Data: models.Message{Message: "Hacı bizi heklemeye çalışma ;)"}})
	}

	if itemInArray(int(user.ID), project.Voters) {
		return c.JSON(http.StatusConflict, models.BaseResponse{Error: true, Data: models.Message{"Zaten bu projeye oy verdin"}})
	}

	if !itemInArray(int(user.ID), project.Voters) {
		userID := int64(user.ID)
		projectID := int64(project.ID)

		project.Voters = append(project.Voters, userID)
		project.VoteCount = int64(len(project.Voters))
		user.Votes = append(user.Votes, projectID)

		err = config.DB.Model(&project).Where("id = ?", project.ID).Update("voters", project.Voters).Update("vote_count", project.VoteCount).Error
		if err != nil {
			return c.JSON(404, err)
		}

		err = config.DB.Model(&user).Where("id = ?", user.ID).Update("votes", user.Votes).Error

		if err != nil {
			return c.JSON(http.StatusNotFound, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
		}

		return c.JSON(200, models.BaseResponse{Data: map[string]string{"message": "Başarıyla oy verdiniz!"}})

	}

	return nil

}

// Index godoc
// @Summary Project Downvote
// @Description Project Downvote
// @Tags Projects
// @Produce  json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 404 {string} string
// @Failure 500 {string} string
// @Router /projects/{id}/ [delete]
// @ID get-string-by-int
// @Param id path int true "Project"
func ProjectDownvote(c echo.Context) error {
	tmpUser := helpers.AuthGetUser(c)
	if tmpUser == nil {
		return c.JSON(http.StatusUnauthorized, models.BaseResponse{Error: true, Data: models.Message{Message: "Oturum açman lazım cano"}})
	}
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
	}

	var project models.Project
	var user models.User

	err = config.DB.Where("id = ?", id).First(&project).Error
	if err != nil {
		return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
	}

	err = config.DB.Where("id = ?", tmpUser.ID).First(&user).Error

	if err != nil {
		return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
	}

	if user.TeamName == project.TeamName {
		return c.JSON(http.StatusTeapot, models.BaseResponse{Error: true, Data: models.Message{Message: "Hacı bizi heklemeye çalışma ;)"}})
	}

	if !itemInArray(int(user.ID), project.Voters) {
		return c.JSON(http.StatusConflict, models.BaseResponse{Error: true, Data: models.Message{Message: "Bu projeye oy vermedin, niye geri almaya çalışıyorsun cano"}})
	}

	if user.Votes == nil {
		return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: "Herhangi bir projeye oy vermedin kral"}})
	}

	if itemInArray(int(user.ID), project.Voters) {
		userID := int64(user.ID)
		projectID := int64(project.ID)

		if len(user.Votes) == 1 {
			user.Votes = nil
		} else {
			user.Votes = removeSelectedIndex(user.Votes, indexOf(projectID, user.Votes))
		}

		if len(project.Voters) == 1 {
			project.Voters = nil
		} else {
			print(indexOf(userID, project.Voters))
			project.Voters = removeSelectedIndex(project.Voters, indexOf(userID, project.Voters))
			print(project.Voters.Value())
		}

		project.VoteCount = int64(len(project.Voters))

		err = config.DB.Model(&user).Where("id = ?", user.ID).Update("votes", user.Votes).Error

		if err != nil {
			return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
		}

		err = config.DB.Model(&project).Where("id = ?", project.ID).Update("voters", project.Voters).Update("vote_count", project.VoteCount).Error

		if err != nil {
			return c.JSON(404, models.BaseResponse{Error: true, Data: models.Message{Message: err.Error()}})
		}
		return c.JSON(200, models.BaseResponse{Data: map[string]string{"message": "Başarıyla oyunuzu geri aldınız!"}})
	}

	return nil

}

func itemInArray(a int, list pq.Int64Array) bool {
	for _, b := range list {
		if b == int64(a) {
			return true
		}
	}
	return false
}

func indexOf(element int64, list pq.Int64Array) int {
	for index, v := range list {
		if element == v {
			return index
		}
	}
	return -1 //not found.
}

func removeSelectedIndex(s pq.Int64Array, index int) pq.Int64Array {
	s[len(s)-1], s[index] = s[index], s[len(s)-1]
	print(removeSelectedIndex)
	return s[:len(s)-1]
}
