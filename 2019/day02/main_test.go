package main

import "testing"
import "reflect"

import "log"

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
	}

	for _, table := range tables {
		processOpcodes(table.intcodes)
		if !reflect.DeepEqual(table.intcodes, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", table.intcodes, table.output)
		}
	}
}

func TestProcessInstruction(t *testing.T) {
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
		result, err := processInstruction(table.instruction, table.numberOne, table.numberTwo)
		if err != nil {
			log.Fatal(err)
		}
		if result != table.result {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.instruction, result, table.result)
		}
	}
}
