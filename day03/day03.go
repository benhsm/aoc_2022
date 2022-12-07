package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	input, err := os.ReadFile("input3.txt")
	if err != nil {
		log.Fatalf("could not read input file: %v", err)
	}
	fmt.Println(sumPriority(input))
	fmt.Println(sumPriority2(input))
}

func sumPriority(s []byte) int {
	s = s[:len(s)-1] // remove trailing newline
	rucksacks := bytes.Split(s, []byte("\n"))
	sum := 0
	for _, r := range rucksacks {
		doubled := make(map[byte]bool)
		seen := make(map[byte]bool)
		for i := 0; i < len(r)/2; i++ {
			seen[r[i]] = true
		}
		for i := len(r) / 2; i < len(r); i++ {
			if seen[r[i]] {
				doubled[r[i]] = true
			}
		}
		for e := range doubled {
			sum += getPriority(e)
		}
	}
	return sum
}

func sumPriority2(s []byte) int {
	s = s[:len(s)-1] // remove trailing newline
	r := bytes.Split(s, []byte("\n"))
	sum := 0
	for i := 0; i < len(r); i += 3 {
		seen := make(map[byte]bool)
		twice := make(map[byte]bool)
		thrice := make(map[byte]bool)
		for _, j := range r[i] {
			seen[j] = true
		}
		for _, k := range r[i+1] {
			if seen[k] {
				twice[k] = true
			}
		}
		for _, l := range r[i+2] {
			if twice[l] {
				thrice[l] = true
			}
		}
		for e := range thrice {
			sum += getPriority(e)
		}
	}
	return sum
}

func getPriority(e byte) int {
	if unicode.IsUpper(rune(e)) {
		return int(e) - 38
	} else if unicode.IsLower(rune(e)) {
		return int(e) - 96
	}
	return 0
}
