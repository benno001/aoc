package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	numbers := parseInput(input)
	result := runPhases(numbers, 100)
	fmt.Println("Answer part 1:", result[0:8])

	realSignal := getRealSignal(input, 10000)
	offset, _ := strconv.Atoi(realSignal[0:7])
	numbers = getPartialSignal(realSignal, offset)
	result2 := runPhasesPartialSum(numbers, 100)
	fmt.Println("Answer part 2:", result2[0:8])
}

func getPartialSignal(input string, offset int) (output []int) {
	// The first half doesn't matter since offset is quite large
	for _, number := range input[offset:] {
		output = append(output, int(number-'0'))
	}
	return output
}

func getRealSignal(input []string, repeats int) string {
	var realSignal string
	for _, line := range input {
		realSignal = strings.Repeat(strings.TrimSpace(line), 10000)
	}
	return realSignal
}

func parseInput(input []string) (numbers []int) {
	for _, line := range input {
		for _, character := range line {
			number, _ := strconv.Atoi(string(character))
			numbers = append(numbers, number)
		}
	}
	return numbers
}

func runPhasesPartialSum(input []int, amount int) []int {
	for i := 0; i < amount; i++ {
		sum := 0
		for j := len(input) - 1; j >= 0; j-- {
			sum += input[j]
			input[j] = sum % 10
		}
	}
	return input
}

func runPhases(input []int, amount int) []int {
	for i := 0; i < amount; i++ {
		input = executePhase(input)
	}
	return input
}

func executePhase(input []int) (output []int) {
	output = make([]int, len(input))
	for i := range input {
		pattern := generatePattern(i + 1)
		out := 0
		for j := 0; j < len(input); j++ {
			if j == 0 {
				p := pattern[1:]
				out += input[j] * p[j%len(p)]
			} else {
				out += input[j] * pattern[(j+1)%len(pattern)]
			}
		}
		str := strconv.Itoa(out)
		number, err := strconv.Atoi(string(str[len(str)-1]))
		if err != nil {
			log.Fatal("Error converting")
		}
		output[i] = number
	}
	return output
}

func generatePattern(position int) []int {
	basePattern := []int{0, 1, 0, -1}
	var pattern []int
	for _, element := range basePattern {
		for i := 0; i < position; i++ {
			pattern = append(pattern, element)
		}
	}
	return pattern
}
