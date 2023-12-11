package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var nodes = map[string][]string{}

func parseNode(line string) {
	if strings.Contains(line, "=") {
		node := line[0:3]
		l := line[7:10]
		r := line[12:15]

		nodes[node] = []string{l, r}
		//fmt.Printf("%s -> [%s, %s]\n", node, l, r)
	}
}

func main() {
	file, err := os.Open("input4.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var directions []string = nil
	for scanner.Scan() {
		line := scanner.Text()
		if directions == nil {
			directions = strings.Split(line, "")
		}

		parseNode(line)
	}

	steps := 0

	// Part 1
	//
	//cur := "AAA"
	//for steps = 0; ; steps++ {
	//	if cur == "ZZZ" {
	//		break
	//	}
	//
	//	var step string
	//	step, directions = directions[0], directions[1:]
	//	directions = append(directions, step)
	//
	//	if step == "L" {
	//		cur = nodes[cur][0]
	//	} else {
	//		cur = nodes[cur][1]
	//	}
	//}
	//
	//fmt.Println(steps)

	// Part 2
	currentNodes := []string{}
	for k, _ := range nodes {
		if strings.HasSuffix(k, "A") {
			currentNodes = append(currentNodes, k)
		}
	}

	fmt.Printf("Starting: %v\n", currentNodes)

	for steps = 1; ; steps++ {
		var step string
		step, directions = directions[0], directions[1:]
		directions = append(directions, step)

		endZ := 0

		for i, n := range currentNodes {
			next := ""
			if step == "L" {
				next = nodes[n][0]
			} else {
				next = nodes[n][1]
			}
			currentNodes[i] = next

			if strings.HasSuffix(next, "Z") {
				endZ++
			}
		}

		if endZ == len(currentNodes) {
			break
		}

		fmt.Println(currentNodes)
		//time.Sleep(1 * time.Second)
	}

	fmt.Println(steps)

}
