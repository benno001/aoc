package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		intcodes, _ := parseProgram(scanner.Text())
		processOpcodes(intcodes)
		fmt.Printf("Part 1: answer is %d\n", intcodes[0])
		for noun := 0; noun < 100; noun++ {
			for verb := 0; verb < 100; verb++ {
				intcodes, _ := parseProgram(scanner.Text())
				intcodes[1] = noun
				intcodes[2] = verb
				processOpcodes(intcodes)
				if intcodes[0] == 19690720 {
					noun := intcodes[1]
					verb := intcodes[2]
					answer := 100*noun + verb
					fmt.Printf("Part 2: noun is %d, verb is %d, answer is %d\n", noun, verb, answer)
				}
			}
		}
	}
}

func parseProgram(input string) ([]int, error) {
	r := csv.NewReader(strings.NewReader(input))
	result, _ := r.ReadAll()
	for _, record := range result {
		var intcodes []int
		for _, code := range record {
			intcode, err := strconv.Atoi(code)
			intcodes = append(intcodes, intcode)
			check(err)
		}
		return intcodes, nil
	}
	return []int{0}, errors.New("Error parsing input")
}

func processOpcodes(intcodes []int) {
	for position := 0; position < len(intcodes); position += 4 {
		instruction := intcodes[position]
		if instruction == 99 {
			break
		}
		numberOne := intcodes[intcodes[position+1]]
		numberTwo := intcodes[intcodes[position+2]]
		out := intcodes[position+3]
		output, err := processInstruction(instruction, numberOne, numberTwo)
		check(err)
		intcodes[out] = output
	}
}

func processInstruction(instruction int, numberOne int, numberTwo int) (int, error) {
	if instruction == 1 {
		return numberOne + numberTwo, nil
	}
	if instruction == 2 {
		return numberOne * numberTwo, nil
	}
	return 0, errors.New("Unknown instruction")
}
