package models

import (
	"time"
)

type Model struct {
	ID        uint64    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	Model
	Username string `sql:"size:64" json:"username" validate:"max=64"`
	Password string `sql:"size:64" json:"password" validate:"max=64"`
	Role     string `sql:"size:32" json:"role" validate:"max=32"`
}
