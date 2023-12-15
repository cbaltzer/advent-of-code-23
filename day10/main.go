package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction uint

const (
	None Direction = iota + 1
	North
	East
	South
	West
)

type Node struct {
	input  Direction
	output Direction
	x      int
	y      int
	symbol string

	prev *Node
	next *Node

	isStart bool
}

func (n Node) Advance() Node {
	var next Node
	switch n.output {
	case North:
		next = n.GetNorth()
	case East:
		next = n.GetEast()
	case South:
		next = n.GetSouth()
	case West:
		next = n.GetWest()
	}

	n.next = &next
	return next
}

func (n Node) GetNorth() Node {
	dx := n.x
	dy := max(0, n.y-1)

	outStr := string(grid[dy][dx])

	var output Direction
	switch outStr {
	case "|":
		output = North
	case "F":
		output = East
	case "7":
		output = West
	default:
		output = None
	}

	next := Node{input: South, output: output, x: dx, y: dy, symbol: outStr, prev: &n, next: nil}
	return next
}

func (n Node) GetEast() Node {
	dx := min(len(grid[0])-1, n.x+1)
	dy := n.y

	outStr := string(grid[dy][dx])
	var output Direction
	switch outStr {
	case "J":
		output = North
	case "-":
		output = East
	case "7":
		output = South
	default:
		output = None
	}

	next := Node{input: West, output: output, x: dx, y: dy, symbol: outStr, prev: &n, next: nil}
	return next
}

func (n Node) GetSouth() Node {
	dx := n.x
	dy := min(len(grid)-1, n.y+1)

	outStr := string(grid[dy][dx])

	var output Direction
	switch outStr {
	case "L":
		output = East
	case "|":
		output = South
	case "J":
		output = West
	default:
		output = None
	}

	next := Node{input: North, output: output, x: dx, y: dy, symbol: outStr, prev: &n, next: nil}
	return next
}

func (n Node) GetWest() Node {
	dx := max(0, n.x-1)
	dy := n.y

	outStr := string(grid[dy][dx])

	var output Direction
	switch outStr {
	case "L":
		output = North
	case "-":
		output = West
	case "F":
		output = South
	default:
		output = None
	}

	next := Node{input: East, output: output, x: dx, y: dy, symbol: outStr, prev: &n, next: nil}
	return next
}

var path = []Node{}
var grid = []string{}
var gridCopy = [][]string{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	var start Node
	var realStart Node
	for y, line := range grid {
		gridCopy = append(gridCopy, strings.Split(line, ""))
		for x, r := range line {
			location := string(r)

			if location == "S" {
				start = Node{input: None, output: None, x: x, y: y, symbol: "S", prev: nil, next: nil, isStart: true}

				if n := string(grid[y-1][x]); n == "7" || n == "|" || n == "F" {
					realStart = start.GetNorth()
				}
				if e := string(grid[y][x+1]); e == "J" || e == "-" || e == "7" {
					realStart = start.GetEast()
				}
				if s := string(grid[y+1][x]); s == "J" || s == "|" || s == "L" {
					realStart = start.GetSouth()
				}
				if w := string(grid[y+1][x]); w == "L" || w == "-" || w == "F" {
					realStart = start.GetWest()
				}
			}

			gridCopy[y][x] = "."
		}
	}

	var curr = realStart

	path = append(path, start)

	i := 1
	area := 0

	for curr.symbol != "S" && i < 100000000 {
		path = append(path, curr)
		i++
		curr = curr.Advance()
	}
	fmt.Println(i / 2)
	fmt.Printf("\n\n\n")

	// Part 2
	for _, n := range path {
		gridCopy[n.y][n.x] = n.symbol
	}

	for y := 0; y < len(gridCopy); y++ {
		for x := 0; x < len(gridCopy[y]); x++ {
			point := gridCopy[y][x]

			inside := 0
			if point == "." {
				for r := 0; r < x; r++ {
					innerPoint := gridCopy[y][r]

					if strings.ContainsAny(innerPoint, "LJ|S") {
						inside++
					}

				}
			}

			if inside%2 != 0 {
				area += 1
				point = "X"
			}

			fmt.Printf("%s", point)
		}
		fmt.Printf("\t%d\n", area)
	}

	fmt.Println(area)

}
