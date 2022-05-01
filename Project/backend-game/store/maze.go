package store

import (
	"github.com/Abunyawa/back_game/domain"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func (s *store) AddMaze(maze *domain.Maze) error {
	query := "INSERT INTO game.maze (map) VALUES ($1) RETURNING id"

	if err := s.db.QueryRow(query, pq.Array(maze.Rows)).Scan(&maze.Id); err != nil {
		log.WithFields(log.Fields{
			"method":  domain.GENERATING,
			"message": err.Error(),
		}).Error("database insert error")
		return err
	}

	log.WithFields(log.Fields{
		"method": domain.GENERATING,
		"id":     maze.Id,
	}).Info("Database insert success")

	return nil
}

func (s *store) GetMaze(id int) (*domain.Maze, error) {
	query := "SELECT id, map FROM game.maze WHERE id = $1"

	ret := &domain.Maze{}
	if err := s.db.QueryRow(query, id).Scan(&ret.Id, pq.Array(&ret.Rows)); err != nil {
		log.WithFields(log.Fields{
			"method":  domain.READING,
			"message": err.Error(),
		}).Error("database query error")
		return nil, err
	}

	log.WithFields(log.Fields{
		"method": domain.GENERATING,
		"id":     ret.Id,
	}).Info("Database query success")

	return ret, nil
}
