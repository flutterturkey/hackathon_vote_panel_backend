package models

import (
	"github.com/lib/pq"
	"time"
)

// User model
type User struct {
	Model
	Name     string
	Email    string
	Password string `json:"-"`
	TeamName string
	Login   *time.Time `json:"-"`
	Logout  *time.Time `json:"-"`
	Votes   pq.Int64Array `json:"-" gorm:"type:integer[]"`
}

const UserVoteLimit = 3