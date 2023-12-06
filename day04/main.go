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
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cardCounts := map[int]int{}

	total := 0
	lineNum := 0
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

		cardCounts[lineNum] += 1 // count the originals
		for c := 0; c < cardCounts[lineNum]; c++ {
			for i := lineNum + 1; i <= lineNum+matchCount; i++ {
				cardCounts[i] += 1
			}
		}

		points := int(math.Pow(2, float64(matchCount-1)))

		if matchCount > 0 {
			total += points
		}

		lineNum += 1
	}

	totalCards := 0
	for _, count := range cardCounts {
		totalCards += count
	}

	fmt.Println(total)
	fmt.Println(totalCards)
}
