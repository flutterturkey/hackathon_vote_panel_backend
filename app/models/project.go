package models

import (
	"github.com/lib/pq"
)

// Project model
type Project struct {
	Model
	Name        string           `json:"name" validate:"required"`
	Description string           `json:"desc" validate:"required"`
	TeamName    string           `json:"team_name" validate:"required"`
	IsLiked 	bool			 `json:"liked" validate:"required"`
	GithubURL   string           `json:"github" validate:"required"`
	VideoURL    string           `json:"video" validate:"required"`
	VoteCount	int64			 `json:"vote"`
	Voters      pq.Int64Array  	 `json:"-" gorm:"type:integer[]"`
	Images      pq.StringArray 	 `json:"images" gorm:"type:text[]"`
}

// Projects model
type Projects struct {
	Model
	Name        string           `json:"name"`
	Description string           `json:"desc"`
	TeamName    string           `json:"team_name"`
	IsLiked 	bool			 `json:"liked" validate:"required"`
	GithubURL   string           `json:"-"`
	VideoURL    string           `json:"-"`
	Voters      pq.Int64Array  	 `json:"-" gorm:"type:integer[]"`
	VoteCount	int64			 `json:"vote"`
	Images      pq.StringArray 	 `json:"-"`
}
