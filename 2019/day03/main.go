package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/juliangruber/go-intersect"
)

type wirePath struct {
	path []string
}

type gridpoint struct {
	x int
	y int
}

type intersectionDistance struct {
	point    gridpoint
	distance int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var wirePaths []wirePath
	for scanner.Scan() {
		wire := parseInput(scanner.Text())[0]
		wirePaths = append(wirePaths, wire)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}

	wireOne := calculatePath(wirePaths[0])
	wireTwo := calculatePath(wirePaths[1])
	intersections := checkIntersection(wireOne, wireTwo)
	origin := gridpoint{x: 0, y: 0}

	var distances []int
	var intersectionSteps []intersectionDistance
	minSteps := 10000000
	for _, val := range intersections {
		distances = append(distances, calculateManhattanDistance(origin, val))
		steps := calculateStepsToIntersection(wireOne, wireTwo, val)
		if steps.distance < minSteps && steps.distance != 0 {
			minSteps = steps.distance
		}
		intersectionSteps = append(intersectionSteps, steps)
	}
	sort.Ints(distances)
	fmt.Printf("Intersections: %v \n", intersections)
	fmt.Printf("Distances: %v \n", distances)
	fmt.Printf("Steps per intersection: %v \n", intersectionSteps)
	fmt.Printf("Min steps %v \n", minSteps)
}

func parseInput(input string) []wirePath {
	r := csv.NewReader(strings.NewReader(input))
	lines, _ := r.ReadAll()
	var wirePaths []wirePath
	for _, record := range lines {
		var wirePath wirePath
		for _, code := range record {
			wirePath.path = append(wirePath.path, code)
		}
		wirePaths = append(wirePaths, wirePath)
	}
	return wirePaths
}

func calculatePath(wire wirePath) []gridpoint {
	gridpoints := []gridpoint{gridpoint{x: 0, y: 0}}
	for _, code := range wire.path {
		gridpoints = append(gridpoints, calculateGridpoint(code, gridpoints[len(gridpoints)-1])...)
	}
	return gridpoints
}

func calculateGridpoint(code string, previousPoint gridpoint) []gridpoint {
	runes := []rune(code)
	instruction := string(runes[0])
	move, err := strconv.Atoi(string(runes[1:]))
	if err != nil {
		log.Fatal("Error converting to integer")
	}
	if instruction == "R" {
		return goRight(previousPoint, move)
	}
	if instruction == "L" {
		return goLeft(previousPoint, move)
	}
	if instruction == "D" {
		return goDown(previousPoint, move)
	}
	if instruction == "U" {
		return goUp(previousPoint, move)
	}
	return []gridpoint{previousPoint}
}

func goLeft(previousPoint gridpoint, move int) []gridpoint {
	var moves []gridpoint
	for i := 1; i <= move; i++ {
		moves = append(moves, gridpoint{x: previousPoint.x - i, y: previousPoint.y})
	}
	return moves
}

func goRight(previousPoint gridpoint, move int) []gridpoint {
	var moves []gridpoint
	for i := 1; i <= move; i++ {
		moves = append(moves, gridpoint{x: previousPoint.x + i, y: previousPoint.y})
	}
	return moves
}

func goUp(previousPoint gridpoint, move int) []gridpoint {
	var moves []gridpoint
	for i := 1; i <= move; i++ {
		moves = append(moves, gridpoint{x: previousPoint.x, y: previousPoint.y + i})
	}
	return moves
}

func goDown(previousPoint gridpoint, move int) []gridpoint {
	var moves []gridpoint
	for i := 1; i <= move; i++ {
		moves = append(moves, gridpoint{x: previousPoint.x, y: previousPoint.y - i})
	}
	return moves
}

func checkIntersection(wireA []gridpoint, wireB []gridpoint) []gridpoint {
	intersection := intersect.Hash(wireA, wireB)
	var gridpoints []gridpoint
	for _, val := range intersection {
		gridpoints = append(gridpoints, val.(gridpoint))
	}
	return gridpoints
}

func calculateManhattanDistance(a gridpoint, b gridpoint) int {
	horizontalDistance := math.Abs(float64(a.x) - float64(b.x))
	verticalDistance := math.Abs(float64(a.y) - float64(b.y))
	return int(horizontalDistance) + int(verticalDistance)
}

func calculateStepsToIntersection(wireA []gridpoint, wireB []gridpoint, intersection gridpoint) intersectionDistance {
	distA := 0
	distB := 0
	for p, v := range wireA {
		if v == intersection {
			distA = p
		}
	}

	for p, v := range wireB {
		if v == intersection {
			distB = p
		}
	}
	return intersectionDistance{intersection, distA + distB}
}
