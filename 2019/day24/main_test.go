package main

import "testing"

import "reflect"

func TestSimulate(t *testing.T) {
	tables := []struct {
		state   [][]string
		minutes int
		result  [][]string
	}{
		{
			[][]string{
				{".", ".", ".", ".", "#"},
				{"#", ".", ".", "#", "."},
				{"#", ".", ".", "#", "#"},
				{".", ".", "#", ".", "."},
				{"#", ".", ".", ".", "."},
			},
			1,
			[][]string{
				{"#", ".", ".", "#", "."},
				{"#", "#", "#", "#", "."},
				{"#", "#", "#", ".", "#"},
				{"#", "#", ".", "#", "#"},
				{".", "#", "#", ".", "."},
			},
		},
		{
			[][]string{
				{".", ".", ".", ".", "#"},
				{"#", ".", ".", "#", "."},
				{"#", ".", ".", "#", "#"},
				{".", ".", "#", ".", "."},
				{"#", ".", ".", ".", "."},
			},
			2,
			[][]string{
				{"#", "#", "#", "#", "#"},
				{".", ".", ".", ".", "#"},
				{".", ".", ".", ".", "#"},
				{".", ".", ".", "#", "."},
				{"#", ".", "#", "#", "#"},
			},
		},
		{
			[][]string{
				{".", ".", ".", ".", "#"},
				{"#", ".", ".", "#", "."},
				{"#", ".", ".", "#", "#"},
				{".", ".", "#", ".", "."},
				{"#", ".", ".", ".", "."},
			},
			3,
			[][]string{
				{"#", ".", ".", ".", "."},
				{"#", "#", "#", "#", "."},
				{".", ".", ".", "#", "#"},
				{"#", ".", "#", "#", "."},
				{".", "#", "#", ".", "#"},
			},
		},
		{
			[][]string{
				{".", ".", ".", ".", "#"},
				{"#", ".", ".", "#", "."},
				{"#", ".", ".", "#", "#"},
				{".", ".", "#", ".", "."},
				{"#", ".", ".", ".", "."},
			},
			4,
			[][]string{
				{"#", "#", "#", "#", "."},
				{".", ".", ".", ".", "#"},
				{"#", "#", ".", ".", "#"},
				{".", ".", ".", ".", "."},
				{"#", "#", ".", ".", "."},
			},
		},
	}
	for _, table := range tables {
		result := simulate(table.state, table.minutes)
		if !reflect.DeepEqual(table.result, result) {
			t.Errorf("Output incorrect for %v minutes. Got %v, wanted %v", table.minutes, result, table.result)
		}
	}
}

func TestFirstToAppearTwice(t *testing.T) {
	tables := []struct {
		state  [][]string
		result [][]string
	}{
		{
			[][]string{
				{".", ".", ".", ".", "#"},
				{"#", ".", ".", "#", "."},
				{"#", ".", ".", "#", "#"},
				{".", ".", "#", ".", "."},
				{"#", ".", ".", ".", "."},
			},
			[][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{"#", ".", ".", ".", "."},
				{".", "#", ".", ".", "."},
			},
		},
	}
	for _, table := range tables {
		_, result := firstToAppearTwice(table.state)
		if !reflect.DeepEqual(table.result, result) {
			t.Errorf("Got %v, wanted %v", result, table.result)
		}
	}
}

func TestGetBioDiversity(t *testing.T) {
	tables := []struct {
		state  [][]string
		result int
	}{
		{
			[][]string{
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{".", ".", ".", ".", "."},
				{"#", ".", ".", ".", "."},
				{".", "#", ".", ".", "."},
			},
			2129920,
		},
	}
	for _, table := range tables {
		result := getBioDiversity(table.state)
		if result != table.result {
			t.Errorf("Got %v, wanted %v", result, table.result)
		}
	}

}
