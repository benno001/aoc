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
		intc := make([]int, 10000)
		intc2 := make([]int, 10000)
		copy(intc[0:], intcodes)
		copy(intc2[0:], intcodes)
		input := 1
		grid := processOpcodes(intc, input)
		visualizeGrid(grid)
		intersections := getIntersections(grid)
		fmt.Println("Answer part 1:", sumAlignmentParameters(intersections))
		input = 2
		grid = processOpcodesPart2(intc2, input)
		visualizeGrid(grid)


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

func visualizeGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func getIntersections(grid [][]string) (intersections [][]int){
	for i, row := range grid {
		for j, col := range row {
			if col == "#" {
				var left string
				var up string
				if j-1 < 0 {
					left = "."
				} else {
					left = grid[i][j-1]
				}
				if i-1 < 0 {
					up = "."
				} else {
					up = grid[i-1][j]
				}
				right := grid[i][j+1]
				down := grid[i+1][j]
				if right == "#" && left == "#" && up == "#" && down == "#" {
					intersections = append(intersections, []int{i, j})
				}
			}
		}
	}
	return intersections
}

func sumAlignmentParameters(intersections [][]int) (sum int) {
	for _, intersection := range intersections {
		sum += calculateAlignmentParameter(intersection)
	}
	return sum
}

func calculateAlignmentParameter(intersection []int) int {
	return intersection[0] * intersection[1]
}

func getMovementCommands() (commands []int) {
	movementCommands := []string{"B","A","B","A","B","C","B","C","A","C"}
	a := []string{"L","10","L","12","R","10"}
	b := []string{"R","6","L","10","R","10","R","10"}
	c := []string{"R","6","L","12","L","10"}

	commands = append(commands, translateCommands(movementCommands)...)
	commands = append(commands, translateCommands(a)...)
	commands = append(commands, translateCommands(b)...)
	commands = append(commands, translateCommands(c)...)
	commands = append(commands, translateCommands([]string{"n"})...)
	return commands
}

func translateCommands(input []string) (output []int) {
	for i, c := range input {
		switch c {
		case "A":
			output = append(output, 65)
		case "B":
			output = append(output, 66)
		case "C":
			output = append(output, 67)
		case "R":
			output = append(output, 82)
		case "L":
			output = append(output, 76)
		case "y":
			output = append(output, 121)
		case "n":
			output = append(output, 110)
		case "6":
			output = append(output, 54)
		case "10":
			output = append(output, []int{49,48}...)
		case "12":
			output = append(output, []int{49,50}...)
		
		}
		if i != len(input) - 1{
			output = append(output, 44)
		}
	}
	output = append(output, 10)
	return output
}

func processOpcodesPart2(intcodes []int, inp int) [][]string {
	relativeBase := 0
	increase := 4
	intcodes[0] = inp

	grid := make([][]string, 48)
	for i := range grid {
		grid[i] = make([]string, 50)
	}

	step := 0
	commands := getMovementCommands()

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
			fmt.Println("Asking input, sending", commands[step])

			putParameter(1, commands[step])
			step++
			increase = 2
		}
		if instruction == 4 {
			output := getParameter(1)
			if output != 94 && output != 35 && output != 46 && output != 10 {
				fmt.Println("Answer part 2:", output)
			}
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
	return grid
}


func processOpcodes(intcodes []int, inp int) [][]string {
	relativeBase := 0
	increase := 4
	intcodes[0] = inp

	grid := make([][]string, 48)
	for i := range grid {
		grid[i] = make([]string, 50)
	}
	row := 0
	col := 0

	step := 0
	commands := getMovementCommands()

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
			fmt.Println("Asking input, sending", commands[step])

			putParameter(1, commands[step])
			step++
			increase = 2
		}
		if instruction == 4 {
			output := getParameter(1)
			if output == 94 {
				grid[row][col] = "^"
				col++
			} else if output == 35 {
				grid[row][col] = "#"
				col++
			} else if output == 46 {
				grid[row][col] = "."
				col++
			} else if output == 10 {
				row++
				col = 0
			}
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
	return grid
}
