package controllers

import (
	"clicks/constants"
	"clicks/helpers"
	"clicks/library/database"

	"clicks/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUserController(ctx *gin.Context) {
	user := models.User{}
	contentType := helpers.GetContentType(ctx)
	ctx.Bind(&user)

	if contentType == constants.APPJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	password := user.Password

	users, err := database.GetUser(user)
	// err := config.DB.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}
	comparePass := helpers.ComparePass([]byte(users.Password), []byte(password))
	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unautorized",
			"message": "Invalid Email/Password",
		})
		return
	}

	// expirationTime := time.Now().Add(24 * time.Hour)
	// claims := &models.Claims{
	// 	Username: user.Username,
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: expirationTime.Unix(),
	// 	},
	// }

	token, err := helpers.GenerateToken(users.ID, users.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    400,
			"message": "Error invalid JWT",
			"status":  "Error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}
