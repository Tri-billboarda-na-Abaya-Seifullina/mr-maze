package store

import (
	"database/sql"
)

type Store interface {
	//AddUser(user *domain.User) error
	//VerifyUser(user *domain.User) (*domain.User, error)
}
type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &store{db: db}
}
