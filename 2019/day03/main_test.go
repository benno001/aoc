package main

import "testing"
import "reflect"

func TestCalculateManhattanDistance(t *testing.T) {
	tables := []struct {
		a        gridpoint
		b        gridpoint
		distance int
	}{
		{gridpoint{x: 2, y: 5}, gridpoint{x: 2, y: 5}, 0},
		{gridpoint{x: 3, y: 5}, gridpoint{x: 2, y: 5}, 1},
		{gridpoint{x: 2, y: 5}, gridpoint{x: 2, y: 10}, 5},
		{gridpoint{x: 3, y: 5}, gridpoint{x: 2, y: 10}, 6},
	}

	for _, table := range tables {
		dist := calculateManhattanDistance(table.a, table.b)
		if dist != table.distance {
			t.Errorf("Distance was incorrect, got: %d, want: %d.", dist, table.distance)
		}
	}
}

func TestCalculatePath(t *testing.T) {
	tables := []struct {
		input  wirePath
		output []gridpoint
	}{
		{wirePath{[]string{"R8", "U5", "L5", "D3"}}, []gridpoint{
			gridpoint{x: 0, y: 0},
			gridpoint{x: 1, y: 0},
			gridpoint{x: 2, y: 0},
			gridpoint{x: 3, y: 0},
			gridpoint{x: 4, y: 0},
			gridpoint{x: 5, y: 0},
			gridpoint{x: 6, y: 0},
			gridpoint{x: 7, y: 0},
			gridpoint{x: 8, y: 0},
			gridpoint{x: 8, y: 1},
			gridpoint{x: 8, y: 2},
			gridpoint{x: 8, y: 3},
			gridpoint{x: 8, y: 4},
			gridpoint{x: 8, y: 5},
			gridpoint{x: 7, y: 5},
			gridpoint{x: 6, y: 5},
			gridpoint{x: 5, y: 5},
			gridpoint{x: 4, y: 5},
			gridpoint{x: 3, y: 5},
			gridpoint{x: 3, y: 4},
			gridpoint{x: 3, y: 3},
			gridpoint{x: 3, y: 2}}},
		{wirePath{[]string{"U7", "R6", "D4", "L4"}}, []gridpoint{
			gridpoint{x: 0, y: 0},
			gridpoint{x: 0, y: 1},
			gridpoint{x: 0, y: 2},
			gridpoint{x: 0, y: 3},
			gridpoint{x: 0, y: 4},
			gridpoint{x: 0, y: 5},
			gridpoint{x: 0, y: 6},
			gridpoint{x: 0, y: 7},
			gridpoint{x: 1, y: 7},
			gridpoint{x: 2, y: 7},
			gridpoint{x: 3, y: 7},
			gridpoint{x: 4, y: 7},
			gridpoint{x: 5, y: 7},
			gridpoint{x: 6, y: 7},
			gridpoint{x: 6, y: 6},
			gridpoint{x: 6, y: 5},
			gridpoint{x: 6, y: 4},
			gridpoint{x: 6, y: 3},
			gridpoint{x: 5, y: 3},
			gridpoint{x: 4, y: 3},
			gridpoint{x: 3, y: 3},
			gridpoint{x: 2, y: 3}}},
	}
	for _, table := range tables {
		result := calculatePath(table.input)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.input, result, table.output)
		}
	}
}

func TestCalculateGridpoint(t *testing.T) {
	tables := []struct {
		instruction       string
		previousGridpoint gridpoint
		output            []gridpoint
	}{
		{"U8", gridpoint{x: 0, y: 0}, []gridpoint{
			gridpoint{x: 0, y: 1},
			gridpoint{x: 0, y: 2},
			gridpoint{x: 0, y: 3},
			gridpoint{x: 0, y: 4},
			gridpoint{x: 0, y: 5},
			gridpoint{x: 0, y: 6},
			gridpoint{x: 0, y: 7},
			gridpoint{x: 0, y: 8}},
		},
		{"R8", gridpoint{x: 0, y: 0}, []gridpoint{
			gridpoint{x: 1, y: 0},
			gridpoint{x: 2, y: 0},
			gridpoint{x: 3, y: 0},
			gridpoint{x: 4, y: 0},
			gridpoint{x: 5, y: 0},
			gridpoint{x: 6, y: 0},
			gridpoint{x: 7, y: 0},
			gridpoint{x: 8, y: 0}},
		},
		{"L8", gridpoint{x: 8, y: 0}, []gridpoint{
			gridpoint{x: 7, y: 0},
			gridpoint{x: 6, y: 0},
			gridpoint{x: 5, y: 0},
			gridpoint{x: 4, y: 0},
			gridpoint{x: 3, y: 0},
			gridpoint{x: 2, y: 0},
			gridpoint{x: 1, y: 0},
			gridpoint{x: 0, y: 0},
		}},
		{"D8", gridpoint{x: 8, y: 5}, []gridpoint{
			gridpoint{x: 8, y: 4},
			gridpoint{x: 8, y: 3},
			gridpoint{x: 8, y: 2},
			gridpoint{x: 8, y: 1},
			gridpoint{x: 8, y: 0},
			gridpoint{x: 8, y: -1},
			gridpoint{x: 8, y: -2},
			gridpoint{x: 8, y: -3},
		}},
	}
	for _, table := range tables {
		result := calculateGridpoint(table.instruction, table.previousGridpoint)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.instruction, result, table.output)
		}
	}
}

func TestParseInput(t *testing.T) {
	tables := []struct {
		input  string
		output []wirePath
	}{
		{"R8,U5,L5,D3\nU7,R6,D4,L4", []wirePath{wirePath{[]string{"R8", "U5", "L5", "D3"}}, wirePath{[]string{"U7", "R6", "D4", "L4"}}}},
	}
	for _, table := range tables {
		result := parseInput(table.input)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.input, result, table.output)
		}
	}
}

func TestCheckIntersection(t *testing.T) {
	tables := []struct {
		wireA  []gridpoint
		wireB  []gridpoint
		output []gridpoint
	}{
		{
			[]gridpoint{
				gridpoint{x: 0, y: 0},
				gridpoint{x: 0, y: 1},
				gridpoint{x: 0, y: 2},
				gridpoint{x: 0, y: 3},
				gridpoint{x: 0, y: 4},
				gridpoint{x: 0, y: 5},
				gridpoint{x: 0, y: 6},
				gridpoint{x: 0, y: 7},
				gridpoint{x: 1, y: 7},
				gridpoint{x: 2, y: 7},
				gridpoint{x: 3, y: 7},
				gridpoint{x: 4, y: 7},
				gridpoint{x: 5, y: 7},
				gridpoint{x: 6, y: 7},
				gridpoint{x: 6, y: 6},
				gridpoint{x: 6, y: 5},
				gridpoint{x: 6, y: 4},
				gridpoint{x: 6, y: 3},
				gridpoint{x: 5, y: 3},
				gridpoint{x: 4, y: 3},
				gridpoint{x: 3, y: 3},
				gridpoint{x: 2, y: 3}},
			[]gridpoint{
				gridpoint{x: 0, y: 0},
				gridpoint{x: 1, y: 0},
				gridpoint{x: 2, y: 0},
				gridpoint{x: 3, y: 0},
				gridpoint{x: 4, y: 0},
				gridpoint{x: 5, y: 0},
				gridpoint{x: 6, y: 0},
				gridpoint{x: 7, y: 0},
				gridpoint{x: 8, y: 0},
				gridpoint{x: 8, y: 1},
				gridpoint{x: 8, y: 2},
				gridpoint{x: 8, y: 3},
				gridpoint{x: 8, y: 4},
				gridpoint{x: 8, y: 5},
				gridpoint{x: 7, y: 5},
				gridpoint{x: 6, y: 5},
				gridpoint{x: 5, y: 5},
				gridpoint{x: 4, y: 5},
				gridpoint{x: 3, y: 5},
				gridpoint{x: 3, y: 4},
				gridpoint{x: 3, y: 3},
				gridpoint{x: 3, y: 2}},
			[]gridpoint{
				gridpoint{x: 0, y: 0},
				gridpoint{x: 6, y: 5},
				gridpoint{x: 3, y: 3},
			},
		},
	}
	for _, table := range tables {
		result := checkIntersection(table.wireA, table.wireB)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", result, table.output)
		}
	}
}

func TestGoLeft(t *testing.T) {
	tables := []struct {
		previousPoint gridpoint
		move          int
		output        []gridpoint
	}{
		{
			gridpoint{x: 6, y: 3},
			4,
			[]gridpoint{
				gridpoint{x: 5, y: 3},
				gridpoint{x: 4, y: 3},
				gridpoint{x: 3, y: 3},
				gridpoint{x: 2, y: 3},
			},
		},
	}
	for _, table := range tables {
		result := goLeft(table.previousPoint, table.move)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", result, table.output)
		}
	}
}

func TestGoRight(t *testing.T) {
	tables := []struct {
		previousPoint gridpoint
		move          int
		output        []gridpoint
	}{
		{
			gridpoint{x: 6, y: 3},
			4,
			[]gridpoint{
				gridpoint{x: 7, y: 3},
				gridpoint{x: 8, y: 3},
				gridpoint{x: 9, y: 3},
				gridpoint{x: 10, y: 3},
			},
		},
	}
	for _, table := range tables {
		result := goRight(table.previousPoint, table.move)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", result, table.output)
		}
	}
}

func TestGoUp(t *testing.T) {
	tables := []struct {
		previousPoint gridpoint
		move          int
		output        []gridpoint
	}{
		{
			gridpoint{x: 6, y: 3},
			4,
			[]gridpoint{
				gridpoint{x: 6, y: 4},
				gridpoint{x: 6, y: 5},
				gridpoint{x: 6, y: 6},
				gridpoint{x: 6, y: 7},
			},
		},
	}
	for _, table := range tables {
		result := goUp(table.previousPoint, table.move)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", result, table.output)
		}
	}
}

func TestGoDown(t *testing.T) {
	tables := []struct {
		previousPoint gridpoint
		move          int
		output        []gridpoint
	}{
		{
			gridpoint{x: 6, y: 3},
			4,
			[]gridpoint{
				gridpoint{x: 6, y: 2},
				gridpoint{x: 6, y: 1},
				gridpoint{x: 6, y: 0},
				gridpoint{x: 6, y: -1},
			},
		},
	}
	for _, table := range tables {
		result := goDown(table.previousPoint, table.move)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", result, table.output)
		}
	}
}

func TestCalculateStepsToIntersection(t *testing.T) {
	tables := []struct {
		wireA        []gridpoint
		wireB        []gridpoint
		intersection gridpoint
		output       intersectionDistance
	}{
		{
			[]gridpoint{
				gridpoint{x: 0, y: 0},
				gridpoint{x: 0, y: 1},
				gridpoint{x: 0, y: 2},
				gridpoint{x: 0, y: 3},
				gridpoint{x: 0, y: 4},
				gridpoint{x: 0, y: 5},
				gridpoint{x: 0, y: 6},
				gridpoint{x: 0, y: 7},
				gridpoint{x: 1, y: 7},
				gridpoint{x: 2, y: 7},
				gridpoint{x: 3, y: 7},
				gridpoint{x: 4, y: 7},
				gridpoint{x: 5, y: 7},
				gridpoint{x: 6, y: 7},
				gridpoint{x: 6, y: 6},
				gridpoint{x: 6, y: 5},
				gridpoint{x: 6, y: 4},
				gridpoint{x: 6, y: 3},
				gridpoint{x: 5, y: 3},
				gridpoint{x: 4, y: 3},
				gridpoint{x: 3, y: 3},
				gridpoint{x: 2, y: 3}},
			[]gridpoint{
				gridpoint{x: 0, y: 0},
				gridpoint{x: 0, y: 1},
				gridpoint{x: 0, y: 2},
				gridpoint{x: 0, y: 3},
				gridpoint{x: 0, y: 4},
				gridpoint{x: 0, y: 5},
				gridpoint{x: 0, y: 6},
				gridpoint{x: 0, y: 7},
				gridpoint{x: 1, y: 7},
				gridpoint{x: 2, y: 7},
				gridpoint{x: 3, y: 7},
				gridpoint{x: 4, y: 7},
				gridpoint{x: 5, y: 7},
				gridpoint{x: 6, y: 7},
				gridpoint{x: 6, y: 6},
				gridpoint{x: 6, y: 5},
				gridpoint{x: 6, y: 4},
				gridpoint{x: 6, y: 3},
				gridpoint{x: 5, y: 3},
				gridpoint{x: 4, y: 3},
				gridpoint{x: 3, y: 3},
				gridpoint{x: 2, y: 3}},
			gridpoint{x: 0, y: 1},
			intersectionDistance{gridpoint{x: 0, y: 1}, 2},
		}, {
			[]gridpoint{
				gridpoint{x: 0, y: 0},
				gridpoint{x: 0, y: 1},
				gridpoint{x: 0, y: 2},
				gridpoint{x: 0, y: 3},
				gridpoint{x: 0, y: 4},
				gridpoint{x: 0, y: 5},
				gridpoint{x: 0, y: 6},
				gridpoint{x: 0, y: 7},
				gridpoint{x: 1, y: 7},
				gridpoint{x: 2, y: 7},
				gridpoint{x: 3, y: 7},
				gridpoint{x: 4, y: 7},
				gridpoint{x: 5, y: 7},
				gridpoint{x: 6, y: 7},
				gridpoint{x: 6, y: 6},
				gridpoint{x: 6, y: 5},
				gridpoint{x: 6, y: 4},
				gridpoint{x: 6, y: 3},
				gridpoint{x: 5, y: 3},
				gridpoint{x: 4, y: 3},
				gridpoint{x: 3, y: 3},
				gridpoint{x: 2, y: 3}},
			[]gridpoint{
				gridpoint{x: 0, y: 0},
				gridpoint{x: 0, y: 1},
				gridpoint{x: 0, y: 2},
				gridpoint{x: 0, y: 3},
				gridpoint{x: 0, y: 4},
				gridpoint{x: 0, y: 5},
				gridpoint{x: 0, y: 6},
				gridpoint{x: 0, y: 7},
				gridpoint{x: 1, y: 7},
				gridpoint{x: 2, y: 7},
				gridpoint{x: 3, y: 7},
				gridpoint{x: 4, y: 7},
				gridpoint{x: 5, y: 7},
				gridpoint{x: 6, y: 7},
				gridpoint{x: 6, y: 6},
				gridpoint{x: 6, y: 5},
				gridpoint{x: 6, y: 4},
				gridpoint{x: 6, y: 3},
				gridpoint{x: 5, y: 3},
				gridpoint{x: 4, y: 3},
				gridpoint{x: 3, y: 3},
				gridpoint{x: 2, y: 3}},
			gridpoint{x: 0, y: 3},
			intersectionDistance{gridpoint{x: 0, y: 3}, 6},
		},
	}
	for _, table := range tables {
		result := calculateStepsToIntersection(table.wireA, table.wireB, table.intersection)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", result, table.output)
		}
	}
}
