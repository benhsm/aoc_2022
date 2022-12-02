package main

import (
	"strings"
	"testing"
)

func TestFindMostCalories(t *testing.T) {
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

	t.Run("basic case", func(t *testing.T) {
		answer := findMostCalories(input)
		if answer != 24000 {
			t.Errorf("want answer 24000, got %v", answer)
		}
	})

	t.Run("find top three", func(t *testing.T) {
		input.Seek(0, 0)
		answer := findTopThree(input)
		want := []int{24000, 11000, 10000}
		for i := range want {
			if answer[i] != want[i] {
				t.Errorf("want answer %v, got %v", want, answer)
				break
			}
		}
	})
}
