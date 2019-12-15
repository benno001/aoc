package main

import (
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
		input    int
		output   []int
	}{
		{[]int{1, 0, 0, 0, 99}, 1, []int{2, 0, 0, 0, 99}},
		{[]int{2, 3, 0, 3, 99}, 1, []int{2, 3, 0, 6, 99}},
		{[]int{2, 4, 4, 5, 99, 0}, 1, []int{2, 4, 4, 5, 99, 9801}},
		{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, 1, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
		{[]int{3, 0, 99}, 1, []int{1, 0, 99}},
		{[]int{5, 1, 3, 99}, 1, []int{5, 1, 3, 99}},
		{[]int{5, 0, 3, 99}, 1, []int{5, 0, 3, 99}},
		{[]int{6, 1, 3, 99}, 1, []int{6, 1, 3, 99}},
		{[]int{6, 0, 3, 99}, 1, []int{6, 0, 3, 99}},
		{[]int{1107, 2, 3, 1, 99}, 1, []int{1107, 1, 3, 1, 99}},
		{[]int{1107, 3, 2, 1, 99}, 1, []int{1107, 0, 2, 1, 99}},
		{[]int{1108, 3, 2, 1, 99}, 1, []int{1108, 0, 2, 1, 99}},
		{[]int{1108, 2, 2, 1, 99}, 1, []int{1108, 1, 2, 1, 99}},
		{[]int{104, 1125899906842624, 99}, 1, []int{104, 1125899906842624, 99}},
		{[]int{1102, 34915192, 34915192, 7, 4, 7, 99, 0}, 1, []int{1102, 34915192, 34915192, 7, 4, 7, 99, 1219070632396864}},
		{[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 1,
			[]int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 16, 1, 0, 0}},
	}
	for _, table := range tables {
		processOpcodes(table.intcodes, table.input)
		if !reflect.DeepEqual(table.intcodes, table.output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", table.intcodes, table.output)
		}
	}
}

func TestModelOxygen(t *testing.T) {
	tables := []struct{
		grid [][]int
		source []int
		minutes int
	}{
		{
			[][]int{
				{0,0,0,0,0,0},
				{0,1,1,0,0,0},
				{0,1,0,1,1,0},
				{0,1,4,1,0,0},
				{0,0,0,0,0,0},
			},
			[]int{3,2}, 4,
		},
	}
	for _, table := range tables {
		m := modelOxygen(table.grid, table.source)
		if m != table.minutes {
			t.Errorf("Output was incorrect, got: %v, want: %v", m, table.minutes)
		}
	}
}