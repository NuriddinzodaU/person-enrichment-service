package models

import (
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name        string `json:"name"`
	Surname     string `json:"surname"`
	Patronymic  string `json:"patronymic,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Nationality string `json:"nationality,omitempty"`
	Age         int    `json:"age,omitempty"`
	ID          int64  `json:"id,omitempty"`
}
