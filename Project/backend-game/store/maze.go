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
