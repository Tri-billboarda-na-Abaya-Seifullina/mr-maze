package store

import (
	"database/sql"
	"github.com/Abunyawa/back_auth/domain"
	log "github.com/sirupsen/logrus"
)

func (s *store) AddUser(user *domain.User) error {
	query := "INSERT INTO auth.users (login, password) VALUES ($1, $2)"

	if _, err := s.db.Exec(query, user.Login, user.Password); err != nil {
		log.WithFields(log.Fields{
			"method":  domain.REGISTER,
			"login":   user.Login,
			"message": err.Error(),
		}).Error("database insert error")
		return err
	}

	log.WithFields(log.Fields{
		"method": domain.REGISTER,
		"login":  user.Login,
	}).Info("database insert success")

	return nil
}

func (s *store) VerifyUser(user *domain.User) (*domain.User, error) {
	query := "SELECT id, login FROM auth.users WHERE login = $1 AND password = $2"

	ret := &domain.User{}
	if err := s.db.QueryRow(query, user.Login, user.Password).Scan(&ret.Id, &ret.Login); err != nil {
		if err != sql.ErrNoRows {
			log.WithFields(log.Fields{
				"method":  domain.AUTH,
				"login":   user.Login,
				"message": err.Error(),
			}).Error("verifying user error")
			return nil, err
		} else {
			log.WithFields(log.Fields{
				"method":  domain.AUTH,
				"login":   user.Login,
				"message": err.Error(),
			}).Error("user not found")
			return nil, err
		}
	}

	log.WithFields(log.Fields{
		"method": domain.AUTH,
		"login":  user.Login,
	}).Info("user found")

	return ret, nil
}
