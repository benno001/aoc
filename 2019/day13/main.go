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

const positionMode, immediateMode, relativeMode = "0", "1", "2"

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		intcodes, _ := parseProgram(scanner.Text())
		intc := make([]int, 100000000)
		intc2 := make([]int, 100000000)
		copy(intc[0:], intcodes)
		copy(intc2[0:], intcodes)
		intc2[0] = 2
		input := 1
		output, _, _ := parseOutput(processOpcodes(intc, input))
		fmt.Println("Answer part 1: ", getNumberOfBlocks(output))
		output, _, score := parseOutput(processOpcodes(intc2, input))
		// visualize(grid)
		fmt.Println("Answer part 2: ", score)
		// input = 2
		// processOpcodes(intc2, input)
	}
}

func getNumberOfBlocks(input [][]int) (total int) {
	for _, v := range input {
		if v[2] == 2 {
			total++
		}
	}
	return total
}

func joyStick(input [][]int) int {
	var ballX int
	var paddleX int
	for _, v := range input {
		for j, val := range v {
			if val == 3 {
				paddleX = j
			}
			if val == 4 {
				ballX = j
			}
		}
	}
	if ballX > paddleX {
		fmt.Println("BallX:", ballX)
		fmt.Println("PaddleX:", paddleX)
		return 1
	}
	if ballX < paddleX {
		fmt.Println("BallX:", ballX)
		fmt.Println("PaddleX:", paddleX)
		return -1
	}
	return 0
}

func getScore(input [][]int) int {
	var score int
	for _, v := range input {
		if v[0] == -1 {
			score = v[2]
		}
	}
	return score 
}

func visualize(grid [][]int){	
	for _, v := range grid {
		fmt.Println(v)
	}
}

func parseOutput(input []int) (output [][]int, grid [][]int, score int) {
	grid = make([][]int, 25)
	for i := range grid {
		grid[i] = make([]int, 45)
	}
	for i := 0; i < len(input); i += 3 {
		triplet := input[i:i+3]
		if triplet[0] > 50 {
			continue
		}
		if triplet[0] == -1 || triplet[1] == -1 {
			score = triplet[2]
			continue
		}
		grid[triplet[1]][triplet[0]] = triplet[2]
		output = append(output, triplet)
	}
	
	return output, grid, score
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

func processOpcodes(intcodes []int, input int) (outputs []int){
	relativeBase := 0
	increase := 4
	for position := 0; position < len(intcodes); position += increase {
		paddedIntCode := fmt.Sprintf("%05d", intcodes[position])
		instruction, _ := strconv.Atoi(paddedIntCode[3:5])

		if instruction == 99 {
			break
		}

		parameterMode := map[int]string{
			1: paddedIntCode[2:3],
			2: paddedIntCode[1:2],
			3: paddedIntCode[0:1],
		}
		getParameter := func(pos int) int {
			switch parameterMode[pos] {
			case positionMode:
				return intcodes[intcodes[position+pos]]
			case immediateMode:
				return intcodes[position+pos]
			case relativeMode:
				return intcodes[relativeBase+intcodes[position+pos]]
			}
			return 0
		}

		putParameter := func(pos int, value int) {
			switch parameterMode[pos] {
			case positionMode:
				intcodes[intcodes[position+pos]] = value
				return
			case relativeMode:
				intcodes[relativeBase+intcodes[position+pos]] = value
				return
			}
		}

		if instruction == 1 || instruction == 2 {
			if instruction == 1 {
				putParameter(3, getParameter(1)+getParameter(2))
			} else if instruction == 2 {
				putParameter(3, getParameter(1)*getParameter(2))
			}
			increase = 4
		}
		if instruction == 3 {
			_, grid, _ := parseOutput(outputs)
			// fmt.Println(position)
			input = joyStick(grid)
			fmt.Println("input:", input)
			putParameter(1, input)
			visualize(grid)
			increase = 2
		}
		if instruction == 4 {
			outputs = append(outputs, getParameter(1))

			increase = 2
		}
		if instruction == 5 || instruction == 6 {
			var newPosition int
			if (instruction == 5 && getParameter(1) != 0) || (instruction == 6 && getParameter(1) == 0) {
				newPosition = getParameter(2)
			} else {
				newPosition = position + 3
			}
			increase = newPosition - position
		}
		if instruction == 7 || instruction == 8 {
			if (instruction == 7 && getParameter(1) < getParameter(2)) || (instruction == 8 && getParameter(1) == getParameter(2)) {
				putParameter(3, 1)
			} else {
				putParameter(3, 0)
			}
			increase = 4
		}
		if instruction == 9 {
			relativeBase = relativeBase + getParameter(1)
			increase = 2
		}
	}
	return outputs
}
