package main

import (
	"testing"
)

func TestSumPriority(t *testing.T) {
	testData := `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

	got := sumPriority([]byte(testData))
	if got != 157 {
		t.Errorf("wanted 157, got %v", got)
	}
}
