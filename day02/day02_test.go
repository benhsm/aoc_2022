package main

import "testing"

func TestRPS(t *testing.T) {
	strategy := []byte(`A Y
B X
C Z
`)
	score := rps(strategy)
	if score != 15 {
		t.Errorf("wanted score 15, got %d", score)
	}
}
