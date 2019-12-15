package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"gopkg.in/karalabe/cookiejar.v2/graph"
	"gopkg.in/karalabe/cookiejar.v2/graph/bfs"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const positionMode, immediateMode, relativeMode = "0", "1", "2"

const north, south, west, east = 1, 2, 3, 4

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
		processOpcodes(intc, input)
		input = 2
		processOpcodes(intc2, input)
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

// getNextMove implements a semi-random walk
func getNextMove(robotStatus int, grid [][]int, currentPosition []int, previousPosition []int) int {
	right := grid[currentPosition[0]][currentPosition[1]+1]
	left := grid[currentPosition[0]][currentPosition[1]-1]
	up := grid[currentPosition[0]+1][currentPosition[1]]
	down := grid[currentPosition[0]-1][currentPosition[1]]
	if right == 3 {
		return east
	}
	if up == 3 {
		return north
	}
	if left == 3 {
		return west
	}
	if down == 3 {
		return south
	}

	return rand.Intn(4) + 1
}

func visualizeGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func getShortestPathBfs(g *graph.Graph, sourceVertice int, destinationVertice int) []int {
	b := bfs.New(g, sourceVertice)
	return b.Path(destinationVertice)
}

func modelOxygen(grid [][]int, source []int) int {
	grid[source[0]][source[1]] = 4
	minutes := 0
	for minute := 0; tilesLeft(grid) > 0; minute++ {
		var changes [][]int
		for i, row := range grid {
			for j, tile := range row {
				if tile == 4 {
					right := grid[i][j+1]
					if right == 1 {
						changes = append(changes, []int{i, j + 1})
					}

					left := grid[i][j-1]
					if left == 1 {
						changes = append(changes, []int{i, j - 1})
					}

					up := grid[i+1][j]
					if up == 1 {
						changes = append(changes, []int{i + 1, j})
					}

					down := grid[i-1][j]
					if down == 1 {
						changes = append(changes, []int{i - 1, j})
					}
				}
			}
		}
		for _, change := range changes {
			grid[change[0]][change[1]] = 4
		}
		minutes++
	}
	return minutes
}

func tilesLeft(grid [][]int) int {
	left := 0
	for _, row := range grid {
		for _, tile := range row {
			if tile == 1 {
				left++
			}
		}
	}
	return left
}

func processOpcodes(intcodes []int, input int) {
	relativeBase := 0
	increase := 4
	area := graph.New(10000)
	grid := make([][]int, 100)
	for i := range grid {
		grid[i] = make([]int, 100)
		for j := range grid[i] {
			grid[i][j] = 3
		}
	}
	currentPosition := []int{50, 50}
	previousPosition := currentPosition
	goal := []int{0, 0}
	sourceVertice := currentPosition[0]*100 + currentPosition[1]
	moveTo := currentPosition
	var outputs []int
	var goalVertice int
	steps := 0
	for position := 0; position < len(intcodes); position += increase {
		paddedIntCode := fmt.Sprintf("%05d", intcodes[position])
		instruction, _ := strconv.Atoi(paddedIntCode[3:5])

		if steps > 1000000 {
			visualizeGrid(grid)
			fmt.Println("Answer part 1:", len(getShortestPathBfs(area, sourceVertice, goalVertice)))
			fmt.Println("Answer part 2:", modelOxygen(grid, goal))
			break
		}
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
			var move int
			if steps == 0 {
				move = north
			} else {
				robotStatus := outputs[len(outputs)-1]
				move = getNextMove(robotStatus, grid, currentPosition, previousPosition)
			}
			if move == north {
				moveTo = []int{currentPosition[0] + 1, currentPosition[1]}
			} else if move == south {
				moveTo = []int{currentPosition[0] - 1, currentPosition[1]}
			} else if move == west {
				moveTo = []int{currentPosition[0], currentPosition[1] - 1}
			} else if move == east {
				moveTo = []int{currentPosition[0], currentPosition[1] + 1}
			}
			putParameter(1, move)
			increase = 2
			steps++
		}
		if instruction == 4 {
			robotStatus := getParameter(1)
			if robotStatus == 2 {
				previousPosition = currentPosition
				currentPosition = moveTo
				grid[currentPosition[0]][currentPosition[1]] = 2
				goal = []int{currentPosition[0], currentPosition[1]}
				// fmt.Println("Goal reached!", currentPosition)
				goalVertice = currentPosition[0]*100 + currentPosition[1]
				area.Connect(currentPosition[0]*100+currentPosition[1], previousPosition[0]*100+previousPosition[1])
				// visualizeGrid(grid)
				// break
			} else if robotStatus == 1 {
				// fmt.Println(grid)
				previousPosition = currentPosition
				currentPosition = moveTo
				grid[currentPosition[0]][currentPosition[1]] = 1
				a := currentPosition[0]*100 + currentPosition[1]
				b := previousPosition[0]*100 + previousPosition[1]
				// fmt.Println("connected", a, b)
				area.Connect(a, b)
			} else if robotStatus == 0 {
				grid[moveTo[0]][moveTo[1]] = 0
			}

			outputs = append(outputs, robotStatus)
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
	fmt.Println(goalVertice)
}
