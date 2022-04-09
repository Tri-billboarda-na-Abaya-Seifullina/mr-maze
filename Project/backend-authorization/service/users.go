package service

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"

	"github.com/Abunyawa/back_auth/domain"
	log "github.com/sirupsen/logrus"
)

var (
	ErrorUnauthorized = errors.New("wrong credentials")
)

func (s *service) AddUser(user *domain.User) error {
	user.Password = hash(user.Password)

	log.WithFields(log.Fields{
		"method": domain.REGISTER,
		"login":  user.Login,
	}).Info("Password hashed")

	return s.Store.AddUser(user)
}

func (s *service) AuthUser(user *domain.User) (string, error) {
	user.Password = hash(user.Password)

	log.WithFields(log.Fields{
		"method": domain.AUTH,
		"login":  user.Login,
	}).Info("Password hashed")

	user, err := s.Store.VerifyUser(user)
	if err != nil {
		return "", ErrorUnauthorized
	}

	expTime := time.Now().Add(time.Hour)

	claims := &domain.Claims{
		Id:    user.Id,
		Login: user.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.SignKey)
	if err != nil {
		log.WithFields(log.Fields{
			"method":  domain.AUTH,
			"login":   user.Login,
			"message": err.Error(),
		}).Error("sign error")
		return "", err
	}

	return tokenString, nil
}

func hash(s string) string {
	hsha256 := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", hsha256)
}

func (s *service) RefreshToken(token *domain.Token) (string, error) {
	claims := &domain.Claims{}

	tkn, err := jwt.ParseWithClaims(token.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return s.SignKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", ErrorUnauthorized
		}
		log.WithFields(log.Fields{
			"method":  domain.REFRESH,
			"token":   token.Token,
			"message": err.Error(),
		}).Error("parsing error")
		return "", err
	}

	log.WithFields(log.Fields{
		"method": domain.REFRESH,
		"token":  token.Token,
	}).Info("Token parsed")

	if !tkn.Valid {
		log.WithFields(log.Fields{
			"method":  domain.REFRESH,
			"token":   token.Token,
			"message": err.Error(),
		}).Error("invalid token")
		return "", ErrorUnauthorized
	}

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		log.WithFields(log.Fields{
			"method":  domain.REFRESH,
			"token":   token.Token,
			"message": "too old token",
		}).Error("too old token")
		return "", ErrorUnauthorized
	}

	expirationTime := time.Now().Add(time.Hour)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(s.SignKey)
	if err != nil {
		log.WithFields(log.Fields{
			"method":  domain.REFRESH,
			"token":   token.Token,
			"message": err.Error(),
		}).Error("sign error")
		return "", ErrorUnauthorized
	}

	log.WithFields(log.Fields{
		"method": domain.REFRESH,
		"token":  token.Token,
	}).Info("Token renewed")

	return tokenString, nil
}
