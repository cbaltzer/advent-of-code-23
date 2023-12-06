package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var times = []int{}
var distances = []int{}

func parseLine(line string, label string, output *[]int) {
	if strings.Contains(line, label) {
		split := strings.Split(line, ":")
		s := strings.Split(split[1], " ")
		for _, v := range s {
			if i, err := strconv.Atoi(v); err == nil {
				*output = append(*output, i)
			}
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		parseLine(line, "Time:", &times)
		parseLine(line, "Distance:", &distances)
	}

	// Part 1
	total := 1
	for i, t := range times {
		wins := 0
		record := distances[i]
		for h := 0; h < t; h++ {
			d := h * (t - h)
			if d > record {
				wins++
			}
		}
		total *= wins
	}

	fmt.Println(total)

	// Part2
	t := joinInts(times)
	record := joinInts(distances)
	wins := 0
	for h := 0; h < t; h++ {
		d := h * (t - h)
		if d > record {
			wins++
		}
	}

	fmt.Println(wins)
}

func joinInts(nums []int) int {
	str := ""
	for _, i := range nums {
		istr := strconv.Itoa(i)
		str += istr
	}
	j, _ := strconv.Atoi(str)
	return j
}
