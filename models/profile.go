package models

import "gorm.io/gorm"

type ProfileResponse struct {
	Code    int       `json:"code"`
	Status  string    `json:"status"`
	Message string    `json:"message"`
	Data    []Profile `json:"data"`
}

type Profile struct {
	gorm.Model
	Username string
	IsLike   bool
	Ispass   bool
}
