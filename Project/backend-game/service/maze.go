package service

import (
	"github.com/Abunyawa/back_game/domain"
)

func (s *service) GenerateMaze(length, width int) (string, error) {
	maze := &domain.Maze{Rows: []string{}}

	row := make([]rune, width*2+1)

	dsu := domain.NewDSU(width)

}
