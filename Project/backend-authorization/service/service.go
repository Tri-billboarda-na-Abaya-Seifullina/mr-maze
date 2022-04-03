package service

import (
	"database/sql"

	"github.com/Abunyawa/back_auth/domain"
	"github.com/Abunyawa/back_auth/store"
)

// Methods declared here
type Service interface {
	ExampleServiceMethod(name string) (string, error)
	AddUser(user *domain.User) error
	AuthUser(user *domain.User) (string, error)
}

type service struct {
	Store   store.Store
	SignKey []byte
}

func NewService(db *sql.DB, key string) Service {

	return &service{
		Store:   store.NewStore(db),
		SignKey: []byte(key),
	}
}
