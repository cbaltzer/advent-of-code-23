package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var numbers = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		first := ""
		first_idx := 0
		last := ""
		last_idx := 0

		line := scanner.Text()

		// part 1
		for i, c := range line {
			if _, err := strconv.Atoi(string(c)); err == nil {
				if first == "" {
					first = string(c)
					first_idx = i
				}
				last = string(c)
				last_idx = i
			}
		}

		// part 2
		for k, v := range numbers {
			if i := strings.Index(line, v); i >= 0 {
				if i < first_idx || first == "" {
					first = fmt.Sprintf("%d", k)
					first_idx = i
				}
			}
			if i := strings.LastIndex(line, v); i >= 0 {
				if i > last_idx || last == "" {
					last = fmt.Sprintf("%d", k)
					last_idx = i
				}
			}

		}

		if lineTotal, err := strconv.Atoi(first + last); err == nil {
			total += lineTotal
		}
	}

	fmt.Println(total)
}
