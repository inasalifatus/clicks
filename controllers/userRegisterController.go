package controllers

import (
	"clicks/constants"
	"clicks/helpers"
	"clicks/library/database"
	"clicks/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserController(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	User := models.User{}

	if contentType == constants.APPJSON {
		ctx.ShouldBindJSON(&User)
	} else {
		ctx.ShouldBind(&User)
	}

	if User.IsPremium != true {
		User.IsPremium = false
	}

	user, err := database.InsertUser(User)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": err.Error(),
		})

	}
	ctx.JSON(http.StatusCreated, map[string]interface{}{
		"status": "success",
		"data":   user,
	})
}
