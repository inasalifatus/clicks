package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	Username       string `json:"username"`
	StandardClaims jwt.StandardClaims
}
