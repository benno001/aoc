package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lower, upper := parseRange(scanner.Text())
		validPasswords := 0
		for i := lower; i <= upper; i++ {
			if hasSixDigits(i) && hasTwoSameAdjacentDigits(i) && increasesInValue(i) && checkBounds(i, lower, upper) {
				validPasswords++
			}
		}
		log.Printf("Valid passwords: %v", validPasswords)
	}
}

func parseRange(r string) (int, int) {
	bounds := strings.Split(r, "-")
	lo, err := strconv.Atoi(bounds[0])
	if err != nil {
		log.Fatal("Error converting lower bound to int")
	}
	hi, err := strconv.Atoi(bounds[1])
	if err != nil {
		log.Fatal("Error converting upper bound to int")
	}
	return lo, hi
}

type matchNumbers struct {
	matches int
	number  int
}

func hasTwoSameAdjacentDigits(password int) bool {
	passwordString := strconv.Itoa(password)
	var previous int
	numbers := make([]matchNumbers, 10)
	for _, v := range passwordString {
		current, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal("Error converting integer")
		}
		if previous == current {
			numbers[current].number = current
			numbers[current].matches++
		}
		previous = current
	}
	isValid := false
	for _, v := range numbers {
		if v.matches == 1 {
			isValid = true
		}
	}
	return isValid
}

func hasSixDigits(password int) bool {
	passwordString := strconv.Itoa(password)
	return len(passwordString) == 6
}

func increasesInValue(password int) bool {
	passwordString := strconv.Itoa(password)
	previous := 0
	for _, v := range passwordString {
		current, err := strconv.Atoi(string(v))
		if err != nil {
			log.Fatal("Error converting integer")
		}
		if previous > current {
			return false
		}
		previous = current
	}
	return true
}

func checkBounds(password int, lower int, upper int) bool {
	return password >= lower && password <= upper
}
