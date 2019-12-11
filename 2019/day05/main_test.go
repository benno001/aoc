package main

import (
	"log"
	"reflect"
	"testing"
)

func TestParseProgram(t *testing.T) {
	tables := []struct {
		input  string
		output []int
	}{
		{"0,1,2,3,4,5,6,7,8,9", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"9,8,7,6,5,4,3,2,1,0", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for _, table := range tables {
		result, _ := parseProgram(table.input)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.input, result, table.output)
		}
	}
}

func TestProcessOpcodes(t *testing.T) {
	tables := []struct {
		intcodes []int
		output   []int
	}{
		{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{[]int{3, 0, 99}, []int{5, 0, 99}},
		{[]int{5, 1, 3, 99}, []int{5, 1, 3, 99}},
		{[]int{5, 0, 3, 99}, []int{5, 0, 3, 99}},
		{[]int{6, 1, 3, 99}, []int{6, 1, 3, 99}},
		{[]int{6, 0, 3, 99}, []int{6, 0, 3, 99}},
		{[]int{1107, 2, 3, 1, 99}, []int{1107, 1, 3, 1, 99}},
		{[]int{1107, 3, 2, 1, 99}, []int{1107, 0, 2, 1, 99}},
		{[]int{1108, 3, 2, 1, 99}, []int{1108, 0, 2, 1, 99}},
		{[]int{1108, 2, 2, 1, 99}, []int{1108, 1, 2, 1, 99}},
	}

	for _, table := range tables {
		processOpcodes(table.intcodes)
		if !reflect.DeepEqual(table.intcodes, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", table.intcodes, table.output)
		}
	}
}

func TestProcessOneTwoInstruction(t *testing.T) {
	tables := []struct {
		instruction int
		numberOne   int
		numberTwo   int
		result      int
	}{
		{1, 1, 1, 2},
		{2, 1, 1, 1},
		{2, 2, 1, 2},
		{1, 3, 7, 10},
		{2, 1, 5, 5},
	}

	for _, table := range tables {
		result, err := processOneTwoInstruction(table.instruction, table.numberOne, table.numberTwo)
		if err != nil {
			log.Fatal(err)
		}
		if result != table.result {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.instruction, result, table.result)
		}
	}
}

func TestProcessOpcodeThree(t *testing.T) {
	tables := []struct {
		intcodes  []int
		numberOne int
		input     int
		result    []int
	}{
		{[]int{3, 3, 2, 1}, 3, 50, []int{3, 3, 2, 50}},
		{[]int{3, 3, 2, 1}, 3, 50, []int{3, 3, 2, 50}},
		{[]int{103, 3, 2, 1}, 3, 50, []int{103, 3, 2, 50}},
	}

	for _, table := range tables {
		processOpcodeThree(table.intcodes, table.numberOne, table.input)
		if !reflect.DeepEqual(table.intcodes, table.result) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", table.intcodes, table.result)
		}
	}
}

func TestProcessOpcodeFour(t *testing.T) {
	tables := []struct {
		intcodes  []int
		numberOne int
		result    int
	}{
		{[]int{4, 3, 2, 1}, 3, 1},
	}

	for _, table := range tables {
		result := processOpcodeFour(table.intcodes, table.numberOne)
		if result != table.result {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.intcodes, result, table.result)
		}
	}
}

func TestProcessOpcodeFiveSix(t *testing.T) {
	tables := []struct {
		instruction int
		position    int
		numberOne   int
		numberTwo   int
		result      int
	}{
		{5, 100, 1, 3, 3},
		{6, 100, 0, 1, 1},
		{5, 100, 0, 1, 103},
		{6, 100, 1, 7, 103},
	}

	for _, table := range tables {
		result := processOpcodeFiveSix(table.instruction, table.numberOne, table.numberTwo, table.position)
		if result != table.result {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.instruction, result, table.result)
		}
	}
}

func TestProcessOpcodeSevenEight(t *testing.T) {
	tables := []struct {
		instruction int
		numberOne   int
		numberTwo   int
		result      int
	}{
		{7, 1, 2, 1},
		{7, 1, 1, 0},
		{8, 1, 1, 1},
		{8, 0, 1, 0},
	}

	for _, table := range tables {
		result := processOpcodeSevenEight(table.instruction, table.numberOne, table.numberTwo)
		if result != table.result {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.instruction, result, table.result)
		}
	}
}

func TestProcessInstruction(t *testing.T) {
	tables := []struct {
		instruction  int
		result       int
		immediateOne bool
		immediateTwo bool
	}{
		{1101, 1, true, true},
		{1001, 1, false, true},
		{101, 1, true, false},
	}

	for _, table := range tables {
		result, immediateOne, immediateTwo := processInstruction(table.instruction)
		if result != table.result || immediateOne != table.immediateOne || immediateTwo != table.immediateTwo {
			t.Errorf("Output for %v was incorrect, got: %v, %v, %v, want: %v, %v, %v.", table.instruction, result, immediateOne, immediateTwo, table.result, table.immediateOne, table.immediateTwo)
		}
	}
}
