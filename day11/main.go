package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

type galaxy struct {
	x, y int
}

var expansionSize = 100 - 1
var universe = [][]string{}

func main() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		cols := strings.Split(line, "")
		universe = append(universe, cols)
	}

	// vertical expansion
	emptyRows := []int{}
	for y, row := range universe {
		if !slices.Contains(row, "#") {
			emptyRows = append(emptyRows, y)
		}
	}

	// horizontal expansion
	emptyCols := []int{}
	for x, _ := range universe[0] {
		hasGalaxy := false
		for y, _ := range universe {
			if universe[y][x] == "#" {
				hasGalaxy = true
			}
		}
		if !hasGalaxy {
			emptyCols = append(emptyCols, x)
		}
	}

	galaxies := []galaxy{}
	for y, row := range universe {
		for x, _ := range row {
			point := universe[y][x]
			if point == "#" {
				galaxies = append(galaxies, galaxy{x: x, y: y})
			}
			fmt.Printf("%s", point)
		}
		fmt.Printf("\n")
	}

	total := 0
	for i, g1 := range galaxies {
		for j, g2 := range galaxies[i+1:] {

			g1he := 0
			g1ve := 0

			g2he := 0
			g2ve := 0

			for _, xe := range emptyCols {
				if xe < g1.x {
					g1he++
				}
				if xe < g2.x {
					g2he++
				}
			}

			for _, ye := range emptyRows {
				if ye < g1.y {
					g1ve++
				}
				if ye < g2.y {
					g2ve++
				}
			}

			g1x := g1.x + (g1he * expansionSize)
			g1y := g1.y + (g1ve * expansionSize)

			g2x := g2.x + (g2he * expansionSize)
			g2y := g2.y + (g2ve * expansionSize)

			absX := max(g1x, g2x) - min(g1x, g2x)
			absY := max(g1y, g2y) - min(g1y, g2y)

			if g1 != g2 {
				distance := absX + absY
				total += distance
				fmt.Printf("G%d (%d,%d) --> G%d (%d,%d) \t [%d]\n", i+1, g1x, g1y, i+1+j+1, g2x, g2y, distance)
			}

		}
	}

	fmt.Println(total)
}
