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
	fmt.Printf("The elf with the most calories has %d calories\n", findMostCalories(inputFile))
	inputFile.Seek(0, 0)
	top3 := findTopThree(inputFile)
	sum := 0
	for _, n := range top3 {
		sum += n
	}
	fmt.Printf("The three elves with the most calories together have %d calories\n", sum)
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

func findTopThree(input io.Reader) []int {
	r := bufio.NewScanner(input)
	var calories int
	top := []int{0, 0, 0}

	for r.Scan() {
		if r.Text() == "" {
			for i := range top {
				if calories > top[i] {
					tmp := top[i]
					top[i] = calories
					calories = tmp
				}
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
	for i := range top {
		if calories > top[i] {
			tmp := top[i]
			top[i] = calories
			calories = tmp
		}
	}
	return top
}
