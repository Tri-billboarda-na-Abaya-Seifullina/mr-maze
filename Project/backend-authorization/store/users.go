package store

import (
	"github.com/Abunyawa/back_auth/domain"
	log "github.com/sirupsen/logrus"
)

func (s *store) AddUser(user *domain.User) error {
	query := "INSERT INTO auth.users (login, password) VALUES ($1, $2)"

	if _, err := s.db.Exec(query, user.Login, user.Password); err != nil {
		log.WithFields(log.Fields{
			"method": domain.REGISTER,
			"login":  user.Login,
		}).Error("database insert error")
		return err
	}

	log.WithFields(log.Fields{
		"method": domain.REGISTER,
		"login":  user.Login,
	}).Info("database insert success")

	return nil
}
