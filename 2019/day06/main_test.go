package main

import (
	"reflect"
	"testing"
)

func TestCreateOrbits(t *testing.T) {
	tables := []struct {
		input  []string
		output map[string]string
	}{
		{[]string{"COM)BBB", "BBB)CCC", "CCC)DDD", "DDD)EEE", "EEE)FFF", "BBB)GGG", "GGG)HHH", "DDD)III", "EEE)JJJ", "JJJ)KKK", "KKK)LLL"}, 
			map[string]string{
			"BBB":"COM",
			"CCC":"BBB",
			"DDD":"CCC",
			"EEE":"DDD",
			"FFF":"EEE",
			"GGG":"BBB",
			"HHH":"GGG",
			"III":"DDD",
			"JJJ":"EEE",
			"KKK":"JJJ",
			"LLL":"KKK",
		}},
		{[]string{"COM)BBB", "BBB)CCC", "CCC)DDD", "DDD)EEE", "EEE)FFF", "BBB)GGG", "GGG)HHH", "DDD)III", "EEE)JJJ", "JJJ)KKK", "KKK)LLL", "KKK)YOU", "III)SAN"},
		map[string]string{
			"BBB":"COM",
			"CCC":"BBB",
			"DDD":"CCC",
			"EEE":"DDD",
			"FFF":"EEE",
			"GGG":"BBB",
			"HHH":"GGG",
			"III":"DDD",
			"JJJ":"EEE",
			"KKK":"JJJ",
			"LLL":"KKK",
			"SAN":"III",
			"YOU":"KKK",
		}},
	}

	for _, table := range tables {
		output := createOrbits(table.input)
		if !reflect.DeepEqual(table.output, output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", output, table.output)
		}
	}
}


func TestCountOrbits(t *testing.T) {
	tables := []struct {
		input map[string]string
		output int
	}{
		{map[string]string{
			"BBB":"COM",
			"CCC":"BBB",
			"DDD":"CCC",
			"EEE":"DDD",
			"FFF":"EEE",
			"GGG":"BBB",
			"HHH":"GGG",
			"III":"DDD",
			"JJJ":"EEE",
			"KKK":"JJJ",
			"LLL":"KKK",
		}, 42},
	}

	for _, table := range tables {
		output := countOrbits(table.input)
		if !reflect.DeepEqual(table.output, output) {
			t.Errorf("Output was incorrect, got: %v, want: %v.", output, table.output)
		}
	}
}