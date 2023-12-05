package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var schematic = []string{}

var gears = map[string][]int{}

type part struct {
	line   int
	start  int
	end    int
	number int
	valid  bool
}

func newPart(lineNum int, line string, number string) part {
	matchxp := fmt.Sprintf(`(^|\D)%s(\D|$)`, number)
	re := regexp.MustCompile(matchxp)

	start := re.FindStringIndex(line)[0] + 1
	end := start + len(number)

	//fmt.Printf("%s - %d:%d\n", number, start, end)
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
		if !hasSymbol {
			hasSymbol = strings.ContainsAny(span, ",/?!@#$%^&*()[]`~;:{}|-=_+ '\\\"")
		}

		if gearIdx := strings.Index(span, "*"); gearIdx >= 0 {
			gearLoc := fmt.Sprintf("%dx%d", minX+gearIdx, i)
			if !slices.Contains(gears[gearLoc], p.number) {
				gears[gearLoc] = append(gears[gearLoc], p.number)
			}
		}

		//fmt.Printf("%d \t %s \t [%d][%d:%d]\n", p.number, span, i, minX, maxX)
	}

	return hasSymbol
}

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
		schematic = append(schematic, scanner.Text())
	}

	for lineNum, line := range schematic {
		re := regexp.MustCompile("\\D")
		numberStrings := re.Split(line, -1)

		for _, n := range numberStrings {
			if n != "" {
				part := newPart(lineNum, line, n)

				if part.checkAdjacent() {
					part.valid = true
					total += part.number
				}
			}
		}
	}

	// gears for pt 2
	gearTotals := 0
	for k, v := range gears {
		if len(v) > 1 {
			ratio := 1
			for _, g := range v {
				ratio *= g
			}
			gearTotals += ratio
			fmt.Printf("[%s] = %v \t\t %d \t\t%d\n", k, v, ratio, gearTotals)
		}

	}

	fmt.Println(total)
	fmt.Println(gearTotals)
}
