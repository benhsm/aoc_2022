package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFile, err := os.OpenFile("input", os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("error reading input data %v", err)
	}
	fmt.Printf("The elf with the most calories has %d calories", findMostCalories(inputFile))
}

func findMostCalories(input io.Reader) int {
	r := bufio.NewScanner(input)
	var maxSeen, calories int
	for r.Scan() {
		if r.Text() == "" {
			if calories > maxSeen {
				maxSeen = calories
			}
			calories = 0
		} else {
			toAdd, err := strconv.Atoi(r.Text())
			if err != nil {
				log.Fatalf("could not interpret data as calories: %v", err)
			}
			calories += toAdd
		}
	}
	return maxSeen
}
