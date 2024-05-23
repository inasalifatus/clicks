package database

import (
	"clicks/config"
	"clicks/models"
)

func GetProfile() (profiles []models.Profile, err error) {
	if err := config.DB.Limit(10).Find(&profiles).Error; err != nil {
		return nil, err
	}
	return profiles, nil
}

func FindProfile(id uint) (profile models.Profile, err error) {
	if err := config.DB.First(&profile).Error; err != nil {
		return profile, err
	}
	return profile, nil
}

func SaveProfile(profile models.Profile) (interface{}, error) {
	if err := config.DB.Save(&profile).Error; err != nil {
		return nil, err
	}
	return profile, nil
}
