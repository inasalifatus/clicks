package helpers

import (
	"clicks/constants"
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GenerateToken(id uint, email string) (signedToken string, err error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(24 * time.Hour).Unix()
	// claims := &models.Claims{
	// 	Username: "",
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	// 	},
	// }

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ = parseToken.SignedString([]byte(constants.SECRET_JWT))

	return signedToken, err
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	errResponse := errors.New("sign in to proceed")
	headerToken := ctx.Request.Header.Get("Authorization")
	bearer := strings.HasPrefix(headerToken, "Bearer")

	if !bearer {
		return nil, errResponse
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errResponse
		}
		return []byte(constants.SECRET_JWT), nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return nil, errResponse
	}
	return token.Claims.(jwt.MapClaims), nil
}
