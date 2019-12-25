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

func main() {
	cards := makeRange(0, 10006)

	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
	deck := runInstructions(cards, input)
	for i, card := range deck {
		if card == 2019 {
			fmt.Println("Answer part 1: ", i)
		}
	}

	largeDeck := 119315717514047
	cardMultiplier := 101741582076661

	// cards2 := makeRange(0, 119315717514047)
	// fmt.Println(cards2[0])
	// for i := 0; i < 10000; i++ {

	// }
}

func runInstructions(deck []int, input []string) (shuffledDeck []int) {
	shuffledDeck = make([]int, len(deck))
	copy(shuffledDeck[0:], deck)
	for _, instruction := range input {
		elements := strings.Split(instruction, " ")
		isCut, _ := regexp.MatchString("cut", elements[0])
		isDeal, _ := regexp.MatchString("deal", elements[0])
		isDealIncr := false
		if len(elements) > 2 {
			isDealIncr, _ = regexp.MatchString("increment", elements[2])
		}
		if isCut {
			index, err := strconv.Atoi(elements[1])
			if err != nil {
				log.Fatal("Error fetching cut index")
			}
			shuffledDeck = cut(shuffledDeck, index)
		}
		if isDeal && !isDealIncr {
			shuffledDeck = deal(shuffledDeck)
		}
		if isDealIncr {
			increment, err := strconv.Atoi(elements[3])
			if err != nil {
				log.Fatal("Error fetching increment")
			}
			shuffledDeck = dealWithIncrement(shuffledDeck, increment)
		}
	}
	return shuffledDeck
}

func cut(deck []int, index int) (shuffledDeck []int) {
	if index < 0 {
		return append(deck[len(deck)+index:], deck[:len(deck)+index]...)
	}
	return append(deck[index:], deck[:index]...)
}

func deal(deck []int) (shuffledDeck []int) {
	shuffledDeck = make([]int, len(deck))
	for i, j := 0, len(deck)-1; i < j; i, j = i+1, j-1 {
		shuffledDeck[i], shuffledDeck[j] = deck[j], deck[i]
	}
	return shuffledDeck
}

func dealWithIncrement(deck []int, increment int) (shuffledDeck []int) {
	shuffledDeck = make([]int, len(deck))
	for i := 0; i < len(deck); i++ {
		shuffledDeck[(i*increment)%len(deck)] = deck[i]
	}
	return shuffledDeck
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func runBigInstructions(cards int, repeats int, input []string) int {
	for _, instruction := range input {
		elements := strings.Split(instruction, " ")
		isCut, _ := regexp.MatchString("cut", elements[0])
		isDeal, _ := regexp.MatchString("deal", elements[0])
		isDealIncr := false
		if len(elements) > 2 {
			isDealIncr, _ = regexp.MatchString("increment", elements[2])
		}
		if isCut {
			index, err := strconv.Atoi(elements[1])
			if err != nil {
				log.Fatal("Error fetching cut index")
			}
			cards = cutBig(shuffledDeck, index)
		}
		if isDeal && !isDealIncr {
			shuffledDeck = dealBig(shuffledDeck)
		}
		if isDealIncr {
			increment, err := strconv.Atoi(elements[3])
			if err != nil {
				log.Fatal("Error fetching increment")
			}
			shuffledDeck = dealWithIncrementBig(shuffledDeck, increment)
		}
	}
	return shuffledDeck
}
