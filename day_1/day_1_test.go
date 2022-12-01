package main

import (
	"strings"
	"testing"
)

func TestFindMostCalories(t *testing.T) {
	t.Run("basic case", func(t *testing.T) {
		input := strings.NewReader(`1000
2000
3000

4000

5000
6000

7000
8000
9000

10000`)
		answer := findMostCalories(input)
		if answer != 24000 {
			t.Errorf("want answer 24000, got %v", answer)
		}
	})
}
