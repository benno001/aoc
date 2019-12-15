package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type reaction struct {
	input  map[string]int
	output map[string]int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
	reactions := parseInput(input)
	ore := 1000000000000
	fmt.Println("Answer part 1:", getOreNeeded(reactions, 1933333))
	fmt.Println("Answer part 2:", getMaxFuel(reactions, ore))
}

func parseInput(input []string) (reactions []reaction) {
	for _, rString := range input {
		elementsIn := make(map[string]int)
		elementsOut := make(map[string]int)
		o := strings.Split(rString, "=> ")
		in := strings.Split(o[0], ", ")
		for _, elIn := range in {
			i := strings.Split(elIn, " ")
			amountIn, _ := strconv.Atoi(i[0])
			elementIn := i[1]
			elementsIn[elementIn] = amountIn
		}
		out := strings.Split(o[1], " ")
		amountOut, _ := strconv.Atoi(out[0])
		elementOut := out[1]
		elementsOut[elementOut] = amountOut
		r := reaction{elementsIn, elementsOut}
		reactions = append(reactions, r)
	}
	return reactions
}

func getOreNeeded(reactions []reaction, fuelNeeded int) (oreNeeded int) {
	// Append all needed items to 'needed' list. Get reaction for needed, add new element.
	// Get list for that element, add to needed, remove that element.
	// If element is 'ORE', get amount needed for producing needed and add to total.

	// First, get elements needed for fuel. Recurse.
	leftOver := make(map[string]int)
	needed := getElementsNeeded(reactions, "FUEL", fuelNeeded, leftOver)
	return needed
}

func getElementsNeeded(reactions []reaction, element string, amount int, leftOver map[string]int) (oreNeeded int) {
	needed := make(map[string]int)
	if element == "ORE" {
		return amount
	}

	for _, r := range reactions {
		if val, ok := r.output[element]; ok {
			if value, excess := leftOver[element]; excess {
				left := value - amount
				if left < 0 {
					amount = int(math.Abs(float64(left)))
					leftOver[element] = 0
				} else {
					leftOver[element] -= amount
					amount = 0
				}
			}
			multiplier := int(math.Ceil(float64(amount) / float64(val)))
			for k, v := range r.input {
				needed[k] = v * multiplier
			}
			overproduction := multiplier*val - amount
			leftOver[element] += overproduction
		}
	}
	ore := 0
	for key, value := range needed {
		ore += getElementsNeeded(reactions, key, value, leftOver)
	}
	return ore
}

func getMaxFuel(reactions []reaction, ore int) int {
	fuel := 0
	start := 0
	end := 10000000
	guesses := 0
	previousGuess := 1
	for {
		fuel = (end-start)/2 + start
		oreRequired := getOreNeeded(reactions, fuel)
		if oreRequired > ore {
			end = fuel
		} else if oreRequired == ore {
			break
		} else {
			start = fuel
		}
		if guesses == 100000 || fuel == previousGuess {
			break
		}
		previousGuess = fuel
		guesses++
	}
	return fuel
}
