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
}

func sumPriority(s []byte) int {
	s = s[:len(s)-1] // remove trailing newline
	rucksacks := bytes.Split(s, []byte("\n"))
	sum := 0
	for num, r := range rucksacks {
		doubled := make(map[byte]bool)
		fmt.Printf("\nrucksack %d\n", num)
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
			var toadd int
			if unicode.IsUpper(rune(e)) {
				toadd = int(e) - 38
			} else if unicode.IsLower(rune(e)) {
				toadd = int(e) - 96
			}
			sum += toadd
			fmt.Printf("%c, %d\n", e, toadd)
		}
	}
	return sum
}
