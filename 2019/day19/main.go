package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"math"
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
		grid := scan(intc)
		// visualizeGrid(grid)
		fmt.Println("Answer part 1:", getTractorBeamTiles(grid))
		x, y := scanShip(intc2)
		
		fmt.Println("Answer part 2:", x*10000+y)
	}
}

func getTractorBeamTiles(grid [][]int) (total int) {
	for _, row := range grid {
		for _, tile := range row {
			if tile == 1 {
				total++
			}
		}
	}
	return total
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

func visualizeGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func calculateShip() (x int, y int) {
	fTop := func(x int) float64 {
		return 0.7274 * float64(x)
	}
	fBottom := func(x int) float64 {
		return 0.9075 * float64(x)
	}
	for x := 100; x < 10000; x++ {
		top := int(math.Round(fTop(x)))
		bottom := int(math.Round(fBottom(x - 99.0)))
		if bottom-top > 99 {
			return x - 99, top
		}
	}
	return 0, 0
}

func scan(intcodes []int) (grid [][]int) {
	grid = make([][]int, 50)
	for i := range grid {
		grid[i] = make([]int, 50)
	}
	for y := 0; y < 50; y++ {
		for x := 0; x < 50; x++ {
			newIntcodes := make([]int, 10000)
			copy(newIntcodes[0:], intcodes)
			input := []int{x, y}
			// fmt.Println(input)
			grid[y][x] = processOpcodes(newIntcodes, input)
		}
	}
	return grid
}

func scanBigArea(intcodes []int) (grid [][]int) {
	grid = make([][]int, 2000)
	for i := range grid {
		grid[i] = make([]int, 2000)
	}
	for y := 500; y < 2000; y++ {
		for x := 500; x < 2000; x++ {
			newIntcodes := make([]int, 10000)
			copy(newIntcodes[0:], intcodes)
			input := []int{x, y}
			// fmt.Println(input)
			grid[y][x] = processOpcodes(newIntcodes, input)
		}
	}
	return grid
}

func scanShip(intcodes []int) (x int, y int) {
	grid := scanBigArea(intcodes)
	for y := 700; y < 2000; y++ {
		for x := 700; x < 2000; x++ {
			if grid[y][x] == 1 && x <= 1901 {
				// fmt.Print(grid[y-99][x+99], "..")
				if grid[y-99][x+99] == 1 {
					return x, y-99
				}
			}
		}
	}
	// visualizePartialGrid(grid, []int{x-200,x+200}, []int{y-200,y+200})
	return 0, 0
}

func processOpcodes(intcodes []int, input []int) (output int) {
	relativeBase := 0
	increase := 4
	firstInput := true
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
			var inp int
			if firstInput {
				inp = input[0]
				firstInput = false
			} else {
				inp = input[1]
			}
			// fmt.Println("instr:", inp)
			putParameter(1, inp)
			increase = 2
		}
		if instruction == 4 {
			output = getParameter(1)
			// fmt.Println("Answer: ", getParameter(1))
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
	return output
}
