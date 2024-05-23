package controllers

import (
	"clicks/library/database"
	"clicks/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ViewProfiles(ctx *gin.Context) {
	user := models.User{}
	userData, err := database.GetUser(user)

	fmt.Println("inas", userData)

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

	profile, err := database.GetProfile()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"status":  "error",
			"message": err.Error(),
		})
	}
	fmt.Println("inas2", profile)
	res := models.ProfileResponse{
		Code:    http.StatusOK,
		Status:  "Success",
		Message: "Success Get Profiles",
		Data:    profile,
	}
	ctx.JSON(http.StatusOK, res)

}
