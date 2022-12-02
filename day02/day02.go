package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

var win = map[byte]byte{
	byte('X'): byte('C'), // rock beats scissors
	byte('Y'): byte('A'), // paper beats rock
	byte('Z'): byte('B'), // scissors beats paper
}

func main() {
	data, err := os.ReadFile("input2.txt")
	if err != nil {
		log.Fatalf("Could not read strategy guide: %v", err)
	}
	fmt.Println("Your score based on the strategy guide is: ", rps(data))
}

func rps(strategy []byte) int {
	strategy = bytes.TrimRight(strategy, "\n") // remove trailing newline
	rounds := bytes.Split(strategy, []byte("\n"))
	var score int
	for _, r := range rounds {
		theirMove, yourMove, fnd := bytes.Cut(r, []byte(" "))
		if !fnd {
			log.Fatalf("could not parse input as round in rps game: %v", r)
		}
		score += calcScore(theirMove[0], yourMove[0])
	}
	return score
}

func calcScore(theirMove, yourMove byte) int {
	score := int(yourMove) - 87
	switch {
	case win[yourMove] == theirMove:
		score += 6
	case int(yourMove)-88 == int(theirMove)-65:
		score += 3
	default:
		score += 0
	}
	return score
}
