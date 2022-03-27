package service

import (
	"crypto/sha256"
	"fmt"

	"github.com/Abunyawa/back_auth/domain"
	log "github.com/sirupsen/logrus"
)

func (s *service) AddUser(user *domain.User) error {
	user.Password = hash(user.Password)

	log.WithFields(log.Fields{
		"method": domain.REGISTER,
		"login":  user.Login,
	}).Info("Password hashed")

	return s.Store.AddUser(user)
}

func hash(s string) string {
	hsha256 := sha256.Sum256([]byte(s))

	return fmt.Sprintf("%x", hsha256)
}
