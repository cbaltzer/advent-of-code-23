package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	partOne()
}

func partOne() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		cards := strings.Split(line, ":")
		numbers := strings.Split(cards[1], "|")

		winning := strings.Split(numbers[0], " ")
		obtained := strings.Split(numbers[1], " ")

		matchCount := 0
		for _, n := range obtained {
			if slices.Contains(winning, n) && n != "" {
				matchCount++
			}
		}

		points := int(math.Pow(2, float64(matchCount-1)))

		if matchCount > 0 {
			total += points
		}
	}

	fmt.Println(total)
}
