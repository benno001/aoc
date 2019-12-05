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
	input := 5
	increase := 4
	instructionMap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
	}
	for position := 0; position < len(intcodes); position += increase {
		instr := intcodes[position]
		instruction, immediateOne, immediateTwo := processInstruction(instr)
		if _, ok := instructionMap[instruction]; !ok || instruction == 99 {
			break
		}
		var numberOne int
		var numberTwo int
		if instruction == 1 || instruction == 2 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			out := intcodes[position+3]
			output, err := processOneTwoInstruction(instruction, numberOne, numberTwo)
			check(err)
			intcodes[out] = output
			increase = 4
		}
		if instruction == 3 {
			numberOne = intcodes[position+1]
			processOpcodeThree(intcodes, numberOne, input)
			increase = 2
		}
		if instruction == 4 {
			numberOne = intcodes[position+1]
			res := processOpcodeFour(intcodes, numberOne)
			fmt.Println("Answer: ", res)
			increase = 2
		}
		if instruction == 5 || instruction == 6 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			newPosition := processOpcodeFiveSix(instruction, numberOne, numberTwo, position)
			increase = newPosition - position
		}
		if instruction == 7 || instruction == 8 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			out := intcodes[position+3]
			output := processOpcodeSevenEight(instruction, numberOne, numberTwo)
			intcodes[out] = output
			increase = 4

		}
	}
}

func processInstruction(instruction int) (int, bool, bool) {
	instr := strconv.Itoa(instruction)
	e, err := strconv.Atoi(string(instr[len(instr) -1]))
	check(err)
	if len(instr) < 3 {
		if e == 9 {
			e2, err := strconv.Atoi(string(instr[len(instr) -2]))
			check(err)
			if e2 == 9 {
				return 99, false, false
			}
		} else {
			return e, false, false
		}
	}
	modeOne, err := strconv.Atoi(string(instr[len(instr) -3]))
	check(err)
	modeTwo := 0
	if len(instr) == 4 {
		modeTwo, err = strconv.Atoi(string(instr[len(instr) -4]))
		check(err)
		return e, modeOne == 1, modeTwo == 1
	}
	if len(instr) > 4 {
		return instruction, false, false
	}
	check(err)
	return e, modeOne == 1, modeTwo == 1
}

func processOneTwoInstruction(instruction int, numberOne int, numberTwo int) (int, error) {
	if instruction == 1 {
		return numberOne + numberTwo, nil
	}
	if instruction == 2 {
		return numberOne * numberTwo, nil
	}
	return 0, errors.New("Unknown instruction")
}

func processOpcodeThree(intcodes []int, param int, input int) {
	intcodes[param] = input
}

func processOpcodeFour(intcodes []int, param int) int {
	return intcodes[param]
}

func processOpcodeFiveSix(instruction int, numberOne int, numberTwo int, position int) int {
	if instruction == 5 {
		if numberOne != 0 {
			return numberTwo
		}
	}
	if instruction == 6 {
		if numberOne == 0 {
			return numberTwo
		}
	}
	return position + 3
}

func processOpcodeSevenEight(instruction int, numberOne int, numberTwo int) int {
	if instruction == 7 {
		if numberOne < numberTwo {
			return 1
		}
	}
	if instruction == 8 {
		if numberOne == numberTwo {
			return 1
		}
	}
	return 0
}
