package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	tables := []struct {
		input  []string
		output []reaction
	}{
		{[]string{"10 ORE => 10 A", "1 ORE => 1 B", "7 A, 1 B => 1 C", "7 A, 1 C => 1 D", "7 A, 1 D => 1 E", "7 A, 1 E => 1 FUEL"},
		[]reaction{
			{map[string]int{
				"ORE": 10,
			}, map[string]int{
				"A": 10,
			}},
			{map[string]int{
				"ORE": 1,
			}, map[string]int{
				"B": 1,
			}},
			{map[string]int{
				"A":7,
				"B":1,
			}, map[string]int{
				"C": 1,
			}}, 
			{map[string]int{
				"A":7,
				"C":1,
			}, map[string]int{
				"D":1,
			}},
			{map[string]int{
				"A":7,
				"D":1,
			}, map[string]int{
				"E":1,
			}},
			{map[string]int{
				"A":7,
				"E":1,
			}, map[string]int{
				"FUEL":1,
			}},
		}},
	}
	for _, table := range tables {
		result := parseInput(table.input)
		if !reflect.DeepEqual(result, table.output) {
			t.Errorf("Output for %v was incorrect, got: %v, want: %v.", table.input, result, table.output)
		}
	}
}

func TestGetOreNeeded(t *testing.T) {
	tables := []struct{
		reactions []reaction
		fuelNeeded int
		oreNeeded int
	}{
		{[]reaction{
			{map[string]int{
				"ORE": 10,
			}, map[string]int{
				"A": 10,
			}},
			{map[string]int{
				"ORE": 1,
			}, map[string]int{
				"B": 1,
			}},
			{map[string]int{
				"A":7,
				"B":1,
			}, map[string]int{
				"C": 1,
			}}, 
			{map[string]int{
				"A":7,
				"C":1,
			}, map[string]int{
				"D":1,
			}},
			{map[string]int{
				"A":7,
				"D":1,
			}, map[string]int{
				"E":1,
			}},
			{map[string]int{
				"A":7,
				"E":1,
			}, map[string]int{
				"FUEL":1,
			}},
		},1,31},
	}
	for _, table := range tables {
		result := getOreNeeded(table.reactions, table.fuelNeeded)
		if result != table.oreNeeded {
			t.Errorf("Output for was incorrect, got: %v, want: %v.", result, table.oreNeeded)
		}
	}
}

func TestGetElementsNeeded(t *testing.T) {
	tables := []struct{
		reactions []reaction
		element string
		amount int
		needed int
	}{
		{
			[]reaction{
				{map[string]int{
					"ORE": 10,
				}, map[string]int{
					"A": 10,
				}},
				{map[string]int{
					"ORE": 1,
				}, map[string]int{
					"B": 1,
				}},
				{map[string]int{
					"A":7,
					"B":1,
				}, map[string]int{
					"C": 1,
				}}, 
				{map[string]int{
					"A":7,
					"C":1,
				}, map[string]int{
					"D":1,
				}},
				{map[string]int{
					"A":7,
					"D":1,
				}, map[string]int{
					"E":1,
				}},
				{map[string]int{
					"A":7,
					"E":1,
				}, map[string]int{
					"FUEL":1,
				}},
			},"A",10,10,
		},
	}
	for _, table := range tables {
		result := getElementsNeeded(table.reactions,table.element,table.amount, make(map[string]int))
		if !reflect.DeepEqual(result, table.needed) {
			t.Errorf("Output for was incorrect, got: %v,  want: %v.", result, table.needed)
		}
	}
}