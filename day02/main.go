package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var colorLimits = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type game struct {
	red   int
	green int
	blue  int
}

func newGame() game {
	g := game{red: 0, green: 0, blue: 0}
	return g
}

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		splitGames := strings.Split(line, ":")
		rolls := splitGames[1]

		game := newGame()

		for _, v := range strings.Split(rolls, ";") {
			for _, color := range strings.Split(v, ",") {
				count, _ := strconv.Atoi(strings.Split(color, " ")[1])

				if strings.Contains(color, "red") && count > game.red {
					game.red = count
				}

				if strings.Contains(color, "green") && count > game.green {
					game.green = count
				}

				if strings.Contains(color, "blue") && count > game.blue {
					game.blue = count
				}
			}
		}

		power := game.red * game.green * game.blue
		total += power
	}

	fmt.Println(total)
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

		valid := true

		splitGames := strings.Split(line, ":")
		label := splitGames[0]
		rolls := splitGames[1]

		for _, v := range strings.Split(rolls, ";") {
			for _, color := range strings.Split(v, ",") {
				colorCount := strings.Split(color, " ")
				count, _ := strconv.Atoi(colorCount[1])
				for colorStr, colorLimit := range colorLimits {
					if strings.Contains(color, colorStr) {
						if count > colorLimit {
							valid = false
						}
					}
				}
			}
		}

		if valid {
			gameID := strings.Split(label, " ")[1]
			if id, err := strconv.Atoi(gameID); err == nil {
				total += id
			}
		}
	}

	fmt.Println(total)
}
