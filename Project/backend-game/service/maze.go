package service

import (
	"github.com/Abunyawa/back_game/domain"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

func (s *service) GenerateMaze(length, width int) (*domain.Map, error) {
	maze := &domain.Maze{Rows: []string{}}

	row := make([]string, width*2+1)

	for i := 0; i < width; i++ {
		row[2*i+1] = "_"
		row[2*i] = "_"
	}
	maze.Rows = append(maze.Rows, strings.Join(row, ""))

	row = make([]string, width*2+1)
	row[0] = "|"
	row[width*2] = "|"
	row[width*2-1] = "_"

	dsu := domain.NewDSU(width)

	for k := 0; k < length; k++ {
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
		open := map[int]int{}
		for k, v := range sz {
			open[k] = rand.Intn(v)
			sz[k] = 0
		}
		for i := 0; i < width; i++ {
			cur := dsu.FindSet(i)
			if open[cur] == sz[cur] {
				row[2*i+1] = " "
			} else {
				res := rand.Intn(10)
				if res > 3 {
					row[2*i+1] = "_"
				} else {
					row[2*i+1] = " "
				}
			}

			sz[cur]++
		}

		if k != length-1 {
			maze.Rows = append(maze.Rows, strings.Join(row, ""))

			newDsu := domain.NewDSU(width)

			sets := map[int]int{}

			for i := 0; i < width; i++ {
				if row[2*i+1] != "_" {
					val, ok := sets[dsu.FindSet(i)]
					if !ok {
						sets[dsu.FindSet(i)] = i
					} else {
						newDsu.UnionSets(val, i)
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
		} else {
			for i := 0; i < width-1; i++ {
				row[2*i+1] = "_"
				if dsu.FindSet(i) != dsu.FindSet(i+1) {
					dsu.UnionSets(i, i+1)
					row[2*i+2] = " "
				}
			}
			maze.Rows = append(maze.Rows, strings.Join(row, ""))
		}
	}

	log.WithFields(log.Fields{
		"method": domain.GENERATING,
		"id":     maze.Id,
	}).Info("Maze generated")

	err := s.Store.AddMaze(maze)
	if err != nil {
		return nil, err
	}

	ret := &domain.Map{Id: maze.Id}
	ret.Cells = make([][]domain.Cell, length)
	for i := 0; i < length; i++ {
		ret.Cells[i] = make([]domain.Cell, width)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			ret.Cells[i][j] = domain.Cell{
				Up:    true,
				Right: true,
				Down:  true,
				Left:  true,
			}
		}
	}

	for i := 0; i <= length; i++ {
		for j, v := range maze.Rows[i] {

			if v == '_' {
				if i != 0 {
					ret.Cells[i-1][j/2].Down = false
				}
				if i < length {
					ret.Cells[i][j/2].Up = false
				}
			}

			if v == '|' {
				if j/2-1 >= 0 {
					ret.Cells[i-1][j/2-1].Right = false
				}
				if j/2 < width {
					ret.Cells[i-1][j/2].Left = false
				}

			}
		}
	}

	return ret, nil
}

func (s *service) GetMaze(id int) (*domain.Map, error) {
	maze, err := s.Store.GetMaze(id)
	length := len(maze.Rows) - 1
	width := len(maze.Rows[0]) / 2

	if err != nil {
		return nil, err
	}

	ret := &domain.Map{Id: maze.Id}
	ret.Cells = make([][]domain.Cell, length)
	for i := 0; i < length; i++ {
		ret.Cells[i] = make([]domain.Cell, width)
	}
	for i := 0; i < length; i++ {
		for j := 0; j < width; j++ {
			ret.Cells[i][j] = domain.Cell{
				Up:    true,
				Right: true,
				Down:  true,
				Left:  true,
			}
		}
	}

	for i := 0; i <= length; i++ {
		for j, v := range maze.Rows[i] {

			if v == '_' {
				if i != 0 {
					ret.Cells[i-1][j/2].Down = false
				}
				if i < length {
					ret.Cells[i][j/2].Up = false
				}
			}

			if v == '|' {
				if j/2-1 >= 0 {
					ret.Cells[i-1][j/2-1].Right = false
				}
				if j/2 < width {
					ret.Cells[i-1][j/2].Left = false
				}

			}
		}
	}

	return ret, nil
}
