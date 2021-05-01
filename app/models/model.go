package models

import "time"

type Model struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time 	`json:"-"`
	UpdatedAt time.Time		`json:"-"`
	DeletedAt *time.Time 	`json:"-" sql:"index"`
}
