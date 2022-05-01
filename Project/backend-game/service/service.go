package service

import (
	"database/sql"
	"github.com/Abunyawa/back_game/domain"

	"github.com/Abunyawa/back_game/store"
)

// Methods declared here
type Service interface {
	GenerateMaze(length, width int) (*domain.Maze, error)
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
