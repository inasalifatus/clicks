package controllers

import (
	"clicks/library/database"
	"clicks/models"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Swipe(ctx *gin.Context) {
	var user models.User
	userData, err := database.GetUser(user)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "error",
		})
		return
	}
	if !userData.IsPremium && userData.SwipeCount >= 10 && userData.LastSwipe.Day() == time.Now().Day() {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error":   "Forbidden",
			"message": "not allowed",
		})
		return
	}

	var input struct {
		ProfileID uint   `json:"profile_id"`
		Action    string `json:"action"`
	}

	if e := ctx.Bind(&input); e != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": e.Error(),
		})
		return
	}

	profile, err := database.FindProfile(input.ProfileID)
	if profile.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"status":  "error",
			"message": "record not found",
		})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "error",
		})
		return
	}

	if input.Action == "like" {
		profile.IsLike = true
	} else if input.Action == "pass" {
		profile.Ispass = true
	} else {
		ctx.JSON(http.StatusForbidden, gin.H{
			"error":   "Forbidden",
			"message": "not allowed",
		})
		return
	}

	saveProfile, err := database.SaveProfile(profile)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	fmt.Println("inas", saveProfile)

	if !userData.IsPremium {
		if userData.LastSwipe.Day() != time.Now().Day() {
			userData.SwipeCount = 0
			userData.LastSwipe = time.Now()
		}
		user.SwipeCount++
		saveUser, err := database.SaveUser(userData)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"status":  "error",
				"message": err.Error(),
			})

		}
		fmt.Println("inas 2", saveUser)
	}

	ctx.JSON(http.StatusOK, models.SwipeResponse{
		Code:    200,
		Status:  "Success",
		Message: "Success update photo",
		Data: models.Swipe{
			Username:  userData.Username,
			IsPremium: userData.IsPremium,
			LastSwipe: userData.LastSwipe,
		},
	})

}
