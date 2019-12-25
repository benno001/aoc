package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"math"
	"os"
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
	grid := parseInput(input)
	minutes, state := firstToAppearTwice(grid)
	fmt.Println("Fist matching state after", minutes, "minutes:")
	visualizeGrid(state)
	fmt.Println("Biodiversity:", getBioDiversity(state))
}

func visualizeGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

func parseInput(input []string) (grid [][]string) {
	grid = make([][]string, len(input))
	for i := range grid {
		grid[i] = make([]string, len(input[0]))
	}
	for i, row := range input {
		for j, char := range row {
			c := string(char)
			grid[i][j] = c
		}
	}
	return grid
}

func simulate(state [][]string, minutes int) (result [][]string) {
	result = make([][]string, len(state))
	for i := range result {
		result[i] = state[i]
	}
	for m := 0; m < minutes; m++ {
		result = step(result)
	}
	return result
}

func step(state [][]string) (result [][]string) {
	result = make([][]string, len(state))
	for i, row := range state {
		result[i] = make([]string, len(row))
		for j, tile := range row {
			up := "."
			if i > 0 {
				up = state[i-1][j]
			}
			down := "."
			if i < len(state)-1 {
				down = state[i+1][j]
			}
			left := "."
			if j > 0 {
				left = state[i][j-1]
			}
			right := "."
			if j < len(state[i])-1 {
				right = state[i][j+1]
			}
			adjacentBugs := 0
			if up == "#" {
				adjacentBugs++
			}
			if down == "#" {
				adjacentBugs++
			}
			if left == "#" {
				adjacentBugs++
			}
			if right == "#" {
				adjacentBugs++
			}
			if tile == "#" {
				if adjacentBugs == 1 {
					result[i][j] = "#"
				} else {
					result[i][j] = "."
				}
			}
			if tile == "." {
				if adjacentBugs == 1 || adjacentBugs == 2 {
					result[i][j] = "#"
				} else {
					result[i][j] = "."
				}
			}
		}
	}
	return result
}

func firstToAppearTwice(state [][]string) (int, [][]string) {
	states := make(map[[32]byte][][]string)
	minutes := 0
	resultState := state
	// Save initial state to the map
	var shaMaterial []byte
	for _, row := range state {
		shaMaterial = append(shaMaterial, []byte(strings.Join(row, ""))...)
	}
	sha256sum := sha256.Sum256(shaMaterial)
	states[sha256sum] = state

	for {
		var shaMaterial []byte
		minutes++
		resultState = step(resultState)
		for _, row := range resultState {
			shaMaterial = append(shaMaterial, []byte(strings.Join(row, ""))...)
		}
		sha256sum := sha256.Sum256(shaMaterial)
		if val, ok := states[sha256sum]; ok {
			resultState = val
			states[sha256sum] = resultState
			break
		} else {
			states[sha256sum] = resultState
		}
	}
	return minutes, resultState
}

func getBioDiversity(state [][]string) (biodiversity int) {
	biodiversity = 0
	for i, row := range state {
		for j, char := range row {
			if char == "#" {
				biodiversity += int(math.Pow(2, float64(i*len(state)+j)))
			}
		}
	}
	return biodiversity
}
