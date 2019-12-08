package main

import (
	"reflect"
	"testing"
)

func TestParseLayer(t *testing.T) {
	tables := []struct {
		input           string
		layerDimensions []int
		output          [][]int
	}{
		{"123456789012", []int{3, 4}, [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {0, 1, 2}}},
		{"0222112222120000", []int{2, 2}, [][]int{{0, 2}, {2, 2}, {1, 1}, {2, 2}, {2, 2}, {1, 2}, {0, 0}, {0, 0}}},
	}
	for _, table := range tables {
		result := parseLayer(table.input, table.layerDimensions)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.input, result, table.output)
		}
	}
}

func TestCountNumbers(t *testing.T) {
	tables := []struct {
		input  [][]int
		result map[int]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {0, 1, 2}}, map[int]int{0: 1, 1: 2, 2: 2, 3: 1, 4: 1, 5: 1, 6: 1, 7: 1, 8: 1, 9: 1}},
	}
	for _, table := range tables {
		result := countNumbers(table.input)
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.input, result, table.result)
		}
	}
}

func TestDecodeImage(t *testing.T) {
	tables := []struct {
		input           map[int][][]int
		layerDimensions []int
		result          [][]int
	}{
		{map[int][][]int{0: {{0, 2}, {2, 2}}, 1: {{1, 1}, {2, 2}}, 2: {{2, 2}, {1, 2}}, 3: {{0, 0}, {0, 0}}}, []int{2, 2}, [][]int{{0, 1}, {1, 0}}},
	}
	for _, table := range tables {
		result := decodeImage(table.input, table.layerDimensions)
		if !reflect.DeepEqual(result, table.result) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.input, result, table.result)
		}
	}
}
