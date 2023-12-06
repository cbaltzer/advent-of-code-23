package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var seeds = []int{}

var seedToSoil = [][]int{}
var soilToFertilizer = [][]int{}
var fertilizerToWater = [][]int{}
var waterToLight = [][]int{}
var lightToTemperature = [][]int{}
var temperatureToHumidity = [][]int{}
var humidityToLocation = [][]int{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var targetMap *[][]int
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "seeds:") {
			split := strings.Split(line, ":")
			s := strings.Split(split[1], " ")
			for _, v := range s {
				if i, err := strconv.Atoi(v); err == nil {
					seeds = append(seeds, i)
				}
			}
		}

		if strings.Contains(line, "seed-to-soil") {
			targetMap = &seedToSoil
		}
		if strings.Contains(line, "soil-to-fertilizer") {
			targetMap = &soilToFertilizer
		}
		if strings.Contains(line, "fertilizer-to-water") {
			targetMap = &fertilizerToWater
		}
		if strings.Contains(line, "water-to-light") {
			targetMap = &waterToLight
		}
		if strings.Contains(line, "light-to-temperature") {
			targetMap = &lightToTemperature
		}
		if strings.Contains(line, "temperature-to-humidity") {
			targetMap = &temperatureToHumidity
		}
		if strings.Contains(line, "humidity-to-location") {
			targetMap = &humidityToLocation
		}

		if targetMap != nil && !strings.Contains(line, ":") {
			values := strings.Split(line, " ")
			mapping := []int{}
			for _, v := range values {
				if i, err := strconv.Atoi(v); err == nil {
					mapping = append(mapping, i)
				}
			}
			if len(mapping) > 0 {
				*targetMap = append(*targetMap, mapping)
			}
		}
	}

	// Part 1
	lowest := -1
	for _, s := range seeds {
		loc := process(s)
		if loc < lowest || lowest == -1 {
			lowest = loc
		}
	}
	fmt.Println(lowest)

	// Part 2
	lowest = -1
	for i := 0; i < len(seeds); i += 2 {
		seedStart := seeds[i]
		seedEnd := seeds[i] + seeds[i+1]

		for s := seedStart; s < seedEnd; s++ {
			loc := process(s)
			if loc < lowest || lowest == -1 {
				lowest = loc
			}
		}
	}
	fmt.Println(lowest)
}

func process(seed int) int {
	a := processMap(seed, seedToSoil)
	b := processMap(a, soilToFertilizer)
	c := processMap(b, fertilizerToWater)
	d := processMap(c, waterToLight)
	e := processMap(d, lightToTemperature)
	f := processMap(e, temperatureToHumidity)
	h := processMap(f, humidityToLocation)
	return h
}

func processMap(input int, mapping [][]int) int {
	for _, v := range mapping {
		srcStart := v[1]
		srcEnd := v[1] + v[2] - 1

		if input >= srcStart && input <= srcEnd {
			//fmt.Println(v)
			delta := input - v[1]
			output := v[0] + delta
			//fmt.Printf("(%d - %d = %d) %d -> %d [%d - %d] \n", input, v[0], delta, input, output, srcStart, srcEnd)
			return output
		}
	}
	return input
}
