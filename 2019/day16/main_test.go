package main

import "testing"

import "reflect"

func TestGeneratePattern(t *testing.T) {
	tables := []struct {
		input  int
		result []int
	}{
		{3, []int{0, 0, 0, 1, 1, 1, 0, 0, 0, -1, -1, -1}},
		{2, []int{0, 0, 1, 1, 0, 0, -1, -1}},
	}
	for _, table := range tables {
		result := generatePattern(table.input)
		if !reflect.DeepEqual(result, table.result){
			t.Errorf("Output was incorrect, got: %v, want: %v", result, table.result)
		}
	}
}

func TestExecutePhase(t *testing.T) {
	tables := []struct {
		input []int
		result []int
	}{
		{[]int{1,2,3,4,5,6,7,8}, []int{4,8,2,2,6,1,5,8}},
	}
	for _, table := range tables {
		result := executePhase(table.input)
		if !reflect.DeepEqual(result, table.result){
			t.Errorf("Output was incorrect, got: %v, want: %v", result, table.result)
		}
	}
}
func TestRunPhases(t *testing.T) {
	tables := []struct {
		input []int
		amount int
		result []int
	}{
		{[]int{1,2,3,4,5,6,7,8}, 1, []int{4,8,2,2,6,1,5,8}},
		{[]int{1,2,3,4,5,6,7,8}, 2, []int{3,4,0,4,0,4,3,8}},
		{[]int{1,2,3,4,5,6,7,8}, 3, []int{0,3,4,1,5,5,1,8}},
		{[]int{1,2,3,4,5,6,7,8}, 4, []int{0,1,0,2,9,4,9,8}},
		{[]int{8,0,8,7,1,2,2,4,5,8,5,9,1,4,5,4,6,6,1,9,0,8,3,2,1,8,6,4,5,5,9,5}, 100, []int{2,4,1, 7, 6, 1, 7, 6, 4, 8, 0, 9, 1, 9, 0, 4, 6, 1 ,1, 4, 0, 3, 8, 7, 6, 3, 1, 9, 5, 5, 9, 5}},
		{[]int{1,9,6,1,7,8,0,4,2,0,7,2,0,2,2,0,9,1,4,4,9,1,6,0,4,4,1,8,9,9,1,7}, 100, []int{7,3, 7, 4, 5, 4, 1, 8, 5, 5, 7, 2, 5, 7, 2, 5, 9, 1, 4, 9, 4, 6, 6, 5, 9, 9, 6, 3, 9, 9, 1, 7}},
		{[]int{6,9,3,1,7,1,6,3,4,9,2,9,4,8,6,0,6,3,3,5,9,9,5,9,2,4,3,1,9,8,7,3}, 100, []int{5, 2, 4, 3, 2, 1, 3, 3, 2, 9, 2, 9, 9, 8, 6, 0, 6, 8, 8, 0, 4, 9, 5, 9, 7, 4, 8, 6, 9, 8, 7, 3}},
	}
	for _, table := range tables {
		result := runPhases(table.input, table.amount)
		if !reflect.DeepEqual(result, table.result){
			t.Errorf("Output was incorrect, got: %v, want: %v", result, table.result)
		}
	}
}