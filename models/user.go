package models

import (
	"clicks/helpers"
	"errors"
	"fmt"
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	ID         uint      `gorm:"primarykey"`
	Username   string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"username" form:"username" valid:"required"`
	Email      string    `gorm:"type:varchar(45);not null;uniqueIndex" json:"email" form:"email" valid:"required"`
	Password   string    `gorm:"type:varchar(100);not null" json:"password" form:"password" valid:"required"`
	Age        int       `json:"age" form:"age" valid:"required"`
	IsPremium  bool      `json:"isPremium"`
	LastSwipe  time.Time `json:"lastSwipe"`
	SwipeCount int       `json:"swipeCount"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	fmt.Println("User Before Create()")

	if u.Age < 8 {
		err = errors.New("age user is too young")
		return
	}

	u.Password = helpers.HashPass(u.Password)
	err = nil
	return
}
