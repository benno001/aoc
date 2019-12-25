package main

import "testing"
import "reflect"

func TestCut(t *testing.T) {
	tables := []struct {
		cards        []int
		index        int
		shuffledDeck []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
			[]int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			-4,
			[]int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
		},
	}
	for _, table := range tables {
		deck := cut(table.cards, table.index)
		if !reflect.DeepEqual(deck, table.shuffledDeck) {
			t.Errorf("Incorrect result for index %v: got %v, want %v", table.index, deck, table.shuffledDeck)
		}
	}
}

func TestDeal(t *testing.T) {
	tables := []struct {
		cards        []int
		shuffledDeck []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, table := range tables {
		deck := deal(table.cards)
		if !reflect.DeepEqual(deck, table.shuffledDeck) {
			t.Errorf("Incorrect result: got %v, want %v", deck, table.shuffledDeck)
		}
	}
}

func TestDealWithIncrement(t *testing.T) {
	tables := []struct {
		cards        []int
		increment    int
		shuffledDeck []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			3,
			[]int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			7,
			[]int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
	}
	for _, table := range tables {
		deck := dealWithIncrement(table.cards, table.increment)
		if !reflect.DeepEqual(deck, table.shuffledDeck) {
			t.Errorf("Incorrect result: got %v, want %v", deck, table.shuffledDeck)
		}
	}
}

func TestRunInstructions(t *testing.T) {
	tables := []struct {
		cards        []int
		instructions []string
		shuffledDeck []int
	}{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]string{"deal with increment 7", "deal into new stack", "deal into new stack"},
			[]int{0, 3, 6, 9, 2, 5, 8, 1, 4, 7},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]string{"cut 6", "deal with increment 7", "deal into new stack"},
			[]int{3, 0, 7, 4, 1, 8, 5, 2, 9, 6},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]string{"deal with increment 7", "deal with increment 9", "cut -2"},
			[]int{6, 3, 0, 7, 4, 1, 8, 5, 2, 9},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]string{
				"deal into new stack",
				"cut -2",
				"deal with increment 7",
				"cut 8",
				"cut -4",
				"deal with increment 7",
				"cut 3",
				"deal with increment 9",
				"deal with increment 3",
				"cut -1",
			},
			[]int{9, 2, 5, 8, 1, 4, 7, 0, 3, 6},
		},
	}
	for _, table := range tables {
		deck := runInstructions(table.cards, table.instructions)
		if !reflect.DeepEqual(deck, table.shuffledDeck) {
			t.Errorf("Incorrect result: got %v, want %v", deck, table.shuffledDeck)
		}
	}
}
