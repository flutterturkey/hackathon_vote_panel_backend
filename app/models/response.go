package models

// Response model
type ProjectsResponse struct {
	Data        []Projects     `json:"data" validate:"required"`
	Error       string         `json:"error" validate:"required"`
}

type ProjectDetailResponse struct {
	Data        Project     `json:"data" validate:"required"`
	Error       string      `json:"error" validate:"required"`
}

type BaseResponse struct {
	Data        interface{}     `json:"data" validate:"required"`
	Error       bool   			`json:"error" validate:"required"`
}

type Message struct {
	Message	string `json:"message" validate:"required"`
}