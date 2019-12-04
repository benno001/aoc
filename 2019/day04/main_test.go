package main

import "testing"

func TestParseRange(t *testing.T) {
	tables := []struct {
		r          string
		lowerBound int
		upperBound int
	}{
		{"111111-222222", 111111, 222222},
	}

	for _, table := range tables {
		lo, up := parseRange(table.r)
		if lo != table.lowerBound || up != table.upperBound {
			t.Errorf("Range was incorrect, got: %d, %d want: %d, %d.", lo, up, table.lowerBound, table.upperBound)
		}
	}
}

func TestCheckBounds(t *testing.T) {
	tables := []struct {
		password   int
		lowerBound int
		upperBound int
		isValid    bool
	}{
		{111112, 111111, 222222, true},
	}

	for _, table := range tables {
		valid := checkBounds(table.password, table.lowerBound, table.upperBound)
		if valid != table.isValid {
			t.Errorf("Distance was incorrect, got: %v want: %v.", valid, table.isValid)
		}
	}
}

func TestHasSixDigits(t *testing.T) {
	tables := []struct {
		password int
		isValid  bool
	}{
		{111112, true},
		{11111, false},
	}

	for _, table := range tables {
		valid := hasSixDigits(table.password)
		if valid != table.isValid {
			t.Errorf("Distance was incorrect, got: %v want: %v.", valid, table.isValid)
		}
	}
}

func TestIncreasesInValue(t *testing.T) {
	tables := []struct {
		password int
		isValid  bool
	}{
		{111112, true},
		{211111, false},
	}

	for _, table := range tables {
		valid := increasesInValue(table.password)
		if valid != table.isValid {
			t.Errorf("Velue was incorrect, got: %v want: %v.", valid, table.isValid)
		}
	}
}

func TestHasTwoSameAdjacentDigits(t *testing.T) {
	tables := []struct {
		password int
		isValid  bool
	}{
		{112233, true},
		{123444, false},
		{111122, true},
	}

	for _, table := range tables {
		valid := hasTwoSameAdjacentDigits(table.password)
		if valid != table.isValid {
			t.Errorf("Value was incorrect, got: %v want: %v.", valid, table.isValid)
		}
	}
}
