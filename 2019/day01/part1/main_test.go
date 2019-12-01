package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{12, 2},
		{14, 2},
		{1969, 654},
		{100756, 33583},
	}

	for _, table := range tables {
		fuel := calculateFuel(table.mass)
		if fuel != table.fuel {
			t.Errorf("Fuel of mass %d was incorrect, got: %d, want: %d.", table.mass, fuel, table.fuel)
		}
	}
}
