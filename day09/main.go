package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input4.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		sequence := strings.Split(line, " ")

		// n/2(2a + (n-1)d)
		n := len(sequence)
		a := sequence[0]

	}

	fmt.Println(total)

}
