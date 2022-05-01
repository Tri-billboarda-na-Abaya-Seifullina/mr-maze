package service

import (
	"github.com/Abunyawa/back_game/domain"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

func (s *service) GenerateMaze(length, width int) (*domain.Maze, error) {
	maze := &domain.Maze{Rows: []string{}}

	row := make([]string, width*2+1)

	for i := 0; i < width; i++ {
		row[2*i+1] = "_"
		row[2*i] = " "
	}
	maze.Rows = append(maze.Rows, strings.Join(row, ""))

	row = make([]string, width*2+1)
	row[0] = "|"
	row[width*2] = "|"
	row[width*2-1] = "_"

	dsu := domain.NewDSU(width)

	for k := 0; k < length-1; k++ {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < width-1; i++ {
			row[2*i+1] = " "
			if dsu.FindSet(i) == dsu.FindSet(i+1) {
				row[2*i+2] = "|"
			} else {
				res := rand.Intn(10)
				if res > 3 {
					row[2*i+2] = " "
					dsu.UnionSets(i, i+1)
				} else {
					row[2*i+2] = "|"
				}
			}
		}

		sz := map[int]int{}
		for i := 0; i < width; i++ {
			sz[dsu.FindSet(i)]++
		}
		masks := map[int]int64{}
		for k, v := range sz {
			maxBitmask := (int64(1) << v) - int64(1)
			sz[k] = 0
			masks[k] = rand.Int63n(maxBitmask) + 1
		}
		for i := 0; i < width; i++ {
			cur := dsu.FindSet(i)

			if (masks[cur] & (1 << sz[cur])) != 0 {
				row[2*i+1] = " "
			} else {
				row[2*i+1] = "_"
			}
			sz[cur]++
		}

		maze.Rows = append(maze.Rows, strings.Join(row, ""))

		newDsu := domain.NewDSU(width)

		for i := 0; i < width-1; i++ {
			if dsu.FindSet(i) == dsu.FindSet(i+1) {
				if row[2*i+1] != "_" && row[2*i+3] != "_" {
					newDsu.UnionSets(i, i+1)
				}
			}
		}

		dsu = newDsu

		for i := 0; i < width*2+1; i++ {
			if row[i] == "_" {
				row[i] = " "
			}
			if row[i] == "|" {
				if k == length-2 {
					continue
				}
				row[i] = " "
			}
		}

		row[0] = "|"
		row[width*2] = "|"
		row[width*2-1] = "_"

	}

	for i := 0; i < width-1; i++ {
		row[2*i+1] = "_"
		if dsu.FindSet(i) != dsu.FindSet(i+1) {
			dsu.UnionSets(i, i+1)
			row[2*i+2] = " "
		}
	}
	maze.Rows = append(maze.Rows, strings.Join(row, ""))

	log.WithFields(log.Fields{
		"method": domain.GENERATING,
		"id":     maze.Id,
	}).Info("Maze generated")

	err := s.Store.AddMaze(maze)
	if err != nil {
		return nil, err
	}

	return maze, nil
}

func (s *service) GetMaze(id int) (*domain.Maze, error) {
	return s.Store.GetMaze(id)
}
