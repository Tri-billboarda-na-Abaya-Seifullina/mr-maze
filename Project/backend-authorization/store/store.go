package store

import (
	"database/sql"

	"github.com/Abunyawa/back_auth/domain"
)

type Store interface {
	AddUser(user *domain.User) error
}
type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &store{db: db}
}
