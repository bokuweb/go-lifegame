package main

import (
	"math/rand"
	"time"
)

type game struct {
	field  [][]bool
	row    int
	column int
}

func newGame(row, column int) *game {
	rand.Seed(time.Now().UnixNano())
	p := new(game)
	p.row = row
	p.column = column
	p.field = make([][]bool, row)

	for r := 0; r < row; r++ {
		p.field[r] = make([]bool, column)
		for c := 0; c < column; c++ {
			if rand.Intn(10) == 0 {
				p.field[r][c] = true
			}
		}
	}
	return p
}

func (p *game) count(r, c int) int {
	if r < 0 || p.row <= r {
		return 0
	}
	if c < 0 || p.column <= c {
		return 0
	}
	if p.field[r][c] {
		return 1
	}
	return 0
}

func (p *game) updateCell(r, c int) bool {
	count := p.count(r-1, c-1) +
		p.count(r-1, c) +
		p.count(r-1, c+1) +
		p.count(r, c-1) +
		p.count(r, c+1) +
		p.count(r+1, c-1) +
		p.count(r+1, c) +
		p.count(r+1, c+1)

	if count == 2 {
		return p.field[r][c]
	} else if count == 3 {
		return true
	} else {
		return false
	}
}

func (p *game) render() {
	print("\033[0;0H")
	for r := 0; r < p.row; r++ {
		for c := 0; c < p.column; c++ {
			cell := " "
			if p.field[r][c] {
				cell = "█"
			}
			print(cell)
		}
		println()
	}
}

func (p *game) update() {
	field := make([][]bool, p.row)
	for r := 0; r < p.row; r++ {
		field[r] = make([]bool, p.column)
		for c := 0; c < p.column; c++ {
			field[r][c] = p.updateCell(r, c)
		}
	}
	p.field = field
}

func main() {
	game := newGame(30, 90)
	game.render()
	ticker := time.Tick(time.Second)
	for {
		<-ticker
		game.update()
		game.render()
	}
}
