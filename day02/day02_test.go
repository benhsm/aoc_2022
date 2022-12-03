package main

import "testing"

func TestRPS(t *testing.T) {
	strategy := []byte(`A Y
B X
C Z
`)
	score := rps(strategy, calcScore)
	if score != 15 {
		t.Errorf("wanted score 15, got %d", score)
	}
}

func TestRPS2(t *testing.T) {
	strategy := []byte(`A Y
B X
C Z
`)
	score := rps(strategy, calcScore2)
	if score != 12 {
		t.Errorf("wanted score 12, got %d", score)
	}
}
