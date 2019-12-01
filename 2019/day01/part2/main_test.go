package main

import "testing"

func TestCalculateFuel(t *testing.T) {
	tables := []struct {
		mass int
		fuel int
	}{
		{14, 2},
		{1969, 966},
		{100756, 50346},
	}

	for _, table := range tables {
		fuel := calculateFuel(table.mass, 0)
		if fuel != table.fuel {
			t.Errorf("Fuel of mass %d was incorrect, got: %d, want: %d.", table.mass, fuel, table.fuel)
		}
	}
}
