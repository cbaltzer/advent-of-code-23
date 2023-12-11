package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type galaxy struct {
	x, y int
}

var universe = [][]string{}
var expandedUniverse = [][]string{}

func main() {
	file, err := os.Open("input.txt")
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
	for _, row := range universe {
		expandedUniverse = append(expandedUniverse, row)
		if !slices.Contains(row, "#") {
			expandedUniverse = append(expandedUniverse, row)
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
	for i, col := range emptyCols {
		for y, _ := range expandedUniverse {
			expandedUniverse[y] = slices.Insert(expandedUniverse[y], col+i, ".")
		}
	}

	galaxies := []galaxy{}
	for y, row := range expandedUniverse {
		for x, _ := range row {
			point := expandedUniverse[y][x]
			if point == "#" {
				galaxies = append(galaxies, galaxy{x: x, y: y})
				point = strconv.Itoa(len(galaxies))
			}

			fmt.Printf("%s", point)
		}
		fmt.Printf("\n")
	}

	total := 0
	for i, g1 := range galaxies {
		for j, g2 := range galaxies[i+1:] {

			absX := max(g1.x, g2.x) - min(g1.x, g2.x)
			absY := max(g1.y, g2.y) - min(g1.y, g2.y)

			if g1 != g2 {
				distance := absX + absY
				total += distance
				fmt.Printf("G%d (%d,%d) --> G%d (%d,%d) \t [%d]\n", i+1, g1.x, g1.y, i+1+j+1, g2.x, g2.y, distance)
			}

		}
	}

	fmt.Println(total)
}
