package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var schematic = []string{}

type part struct {
	line   int
	start  int
	end    int
	number int
	valid  bool
}

func newPart(lineNum int, line string, number string) part {
	start := strings.Index(line, number)
	end := start + len(number)

	fmt.Printf("%s - %d:%d\n", number, start, end)
	value, err := strconv.Atoi(number)
	if err != nil {
		fmt.Println(err)
	}
	_part := part{line: lineNum, start: start, end: end, number: value}

	return _part
}

func (p *part) toString() string {
	if p.valid {
		return fmt.Sprintf("%d", p.number)
	} else {
		return fmt.Sprintf("%d*", p.number)
	}
}

func (p *part) checkAdjacent() bool {
	minY := max(0, p.line-1)
	maxY := min(len(schematic)-1, p.line+1)

	minX := max(0, p.start-1)
	maxX := min(len(schematic[0]), p.end+1)

	hasSymbol := false
	for i := minY; i <= maxY; i++ {
		span := schematic[i][minX:maxX]
		hasSymbol = strings.ContainsAny(span, ",/?!@#$%^&*()[]`~;:{}|-=_+ '\\\"")

		fmt.Printf("%d \t %s \t [%d][%d:%d] \t %d\n", p.number, span, i, minX, maxX, hasSymbol)
	}

	fmt.Println()

	return hasSymbol
}

func main() {
	partOne()
}

func partOne() {
	file, err := os.Open("input2.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}

	for lineNum, line := range schematic {
		//fmt.Printf("%d\t%s\n", lineNum, line)

		re := regexp.MustCompile("\\D")
		numberStrings := re.Split(line, -1)

		for _, n := range numberStrings {
			if n != "" {
				//fmt.Println(n)

				part := newPart(lineNum, line, n)

				if part.checkAdjacent() {
					//fmt.Println(part.number)
					part.valid = true
					total += part.number
				}

				//fmt.Println(part.toString())
			}
		}

	}

	fmt.Println(total)
}
