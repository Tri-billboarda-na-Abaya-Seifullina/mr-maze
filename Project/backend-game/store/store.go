package store

import (
	"database/sql"
	"github.com/Abunyawa/back_game/domain"
)

type Store interface {
	//AddUser(user *domain.User) error
	//VerifyUser(user *domain.User) (*domain.User, error)
	AddMaze(maze *domain.Maze) error
	GetMaze(id int) (*domain.Maze, error)
}
type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &store{db: db}
}
