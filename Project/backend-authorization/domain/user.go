package domain

import (
	"github.com/dgrijalva/jwt-go"
)

const (
	REGISTER = "register"
	AUTH     = "auth"
	REFRESH  = "refresh"
)

type User struct {
	Login    string
	Password string
	Id       int
}

type Claims struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
	jwt.StandardClaims
}

type Token struct {
	Token string
}
