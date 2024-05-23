package database

import (
	"clicks/config"
	"clicks/models"
)

func InsertUser(user models.User) (interface{}, error) {
	if err := config.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func GetUser(user models.User) (out models.User, err error) {
	if err := config.DB.Where("username = ?", user.Username).First(&out).Error; err != nil {
		return out, err
	}
	return out, nil
}

func SaveUser(user models.User) (interface{}, error) {
	if err := config.DB.Save(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
