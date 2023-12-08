package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard HandType = iota + 1
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

const CardPriority = "23456789TJQKA" //"AKQJT98765432"

type Hand struct {
	cards    string
	strength HandType
	bid      int
}

func newHand(line string) Hand {
	s := strings.Split(line, " ")
	hand := s[0]
	bid, _ := strconv.Atoi(s[1])

	cardCounts := map[string]int{}

	for _, r := range hand {
		c := string(r)
		cardCounts[c] += 1
	}

	var strength HandType = HighCard
	for _, v := range cardCounts {
		if v == 5 {
			strength = FiveOfAKind
		}
		if v == 4 {
			strength = FourOfAKind
		}
		if v == 3 {
			if strength == OnePair {
				strength = FullHouse
			} else {
				strength = ThreeOfAKind
			}
		}
		if v == 2 {
			if strength == OnePair {
				strength = TwoPair
			} else if strength == ThreeOfAKind {
				strength = FullHouse
			} else {
				strength = OnePair
			}
		}
	}

	return Hand{cards: hand, strength: strength, bid: bid}
}

func handLess(h1 Hand, h2 Hand) bool {
	if h1.strength == h2.strength {
		for i, r := range h1.cards {
			card1Strength := strings.Index(CardPriority, string(r))
			card2Strength := strings.Index(CardPriority, string(h2.cards[i]))
			if card1Strength != card2Strength {
				return card1Strength < card2Strength
			}
		}
	}
	return h1.strength < h2.strength
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	hands := []Hand{}

	for scanner.Scan() {
		line := scanner.Text()

		hand := newHand(line)
		hands = append(hands, hand)
	}

	sort.Slice(hands, func(i, j int) bool {
		return handLess(hands[i], hands[j])
	})

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
		fmt.Printf("%s %d \t\t[%d * $%d]\n", h.cards, h.strength, i+1, h.bid)
	}

	fmt.Println(total)
}
