package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var doneErr = errors.New("no more moves left")

const (
	Up = iota
	Down
	Left
	Right
)

type Point struct{ x, y int }
type ropeModel struct {
	head      Point
	tail      Point
	times     int
	moves     []string
	grid      map[Point]struct{}
	moveIndex int
	currDir   int
}

func newModel(moves []string) ropeModel {
	return ropeModel{
		head:      Point{0, 0},
		tail:      Point{0, 0},
		times:     0,
		moves:     moves,
		grid:      make(map[Point]struct{}),
		moveIndex: 0,
		currDir:   0,
	}
}

func (m *ropeModel) step() error {
	if m.times == 0 {
		if m.moveIndex+1 > len(m.moves) {
			return doneErr
		}
		m.nextMove()
	}
	switch m.currDir {
	case Up:
		m.head.y++
	case Down:
		m.head.y--
	case Left:
		m.head.x--
	case Right:
		m.head.x++
	}

	switch m.head.y - m.tail.y {
	case 2:
		m.tail.y++
		m.tail.x = m.head.x
	case 1:
		if m.head.x > m.tail.x+1 {
			m.tail.y++
			m.tail.x++
		} else if m.head.x < m.tail.x-1 {
			m.tail.y++
			m.tail.x--
		}
	case 0:
		if m.head.x > m.tail.x+1 {
			m.tail.x++
		} else if m.head.x < m.tail.x-1 {
			m.tail.x--
		}
	case -1:
		if m.head.x > m.tail.x+1 {
			m.tail.y--
			m.tail.x++
		} else if m.head.x < m.tail.x-1 {
			m.tail.y--
			m.tail.x--
		}
	case -2:
		m.tail.y--
		m.tail.x = m.head.x
	}
	m.grid[m.tail] = struct{}{}
	m.times--

	return nil
}

func (m *ropeModel) nextMove() {
	newDir, newTimes, ok := strings.Cut(m.moves[m.moveIndex], " ")
	if !ok {
		log.Fatal("couldn't parse moves!")
	}
	switch newDir {
	case "U":
		m.currDir = Up
	case "D":
		m.currDir = Down
	case "L":
		m.currDir = Left
	case "R":
		m.currDir = Right
	default:
		log.Fatal("couldn't parse moves! got: ", newDir)
	}
	times, err := strconv.Atoi(newTimes)
	if err != nil {
		log.Fatal("couldn't parse moves! err: ", err)
	}
	m.times = times
	m.moveIndex++
}

func (m *ropeModel) solve() int {
	var res int
	for {
		err := m.step()
		if err == doneErr {
			for _ = range m.grid {
				res++
			}
			break
		}
	}
	return res
}

func main() {
	input, err := os.ReadFile("input9.txt")
	if err != nil {
		log.Fatalf("could not read file")
	}
	moves := strings.Split(string(input), "\n")
	moves = moves[:len(moves)-1]
	p1 := newModel(moves)
	answer := p1.solve()
	fmt.Println(answer)
}
