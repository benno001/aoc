package main

import "testing"
import "reflect"

// type asteroid struct {
// 	x float64
// 	y float64
// }

// type path struct {
// 	direction float64
// 	distance  float64
// }

func TestParseMap(t *testing.T) {
	tables := []struct {
		input  []string
		result []asteroid
	}{
		{[]string{
			".#..#",
			".....",
			"#####",
			"....#",
			"...##"},
			[]asteroid{
				{1.5, 0.5}, {4.5, 0.5},
				{0.5, 2.5}, {1.5, 2.5}, {2.5, 2.5}, {3.5, 2.5}, {4.5, 2.5},
				{4.5, 3.5},
				{3.5, 4.5}, {4.5, 4.5},
			},
		},
	}
	for _, table := range tables {
		result := parseMap(table.input)
		if !reflect.DeepEqual(table.result, result) {
			t.Errorf("Map was incorrect, got: %v, want: %v.", result, table.result)
		}
	}
}

func TestGetLine(t *testing.T) {
	tables := []struct {
		a               asteroid
		b               asteroid
		resultDirection float64
		resultDistance  float64
	}{
		{asteroid{1.5, 0.5}, asteroid{4.5, 0.5}, 0, 3},
		{asteroid{0.5, 2.5}, asteroid{4.5, 0.5}, -0.4636476090008061, 4.47213595499958},
	}
	for _, table := range tables {
		resultDirection, resultDistance := getLine(table.a, table.b)
		if resultDirection != table.resultDirection || resultDistance != table.resultDistance {
			t.Errorf("Output was incorrect, got: %v, %v, want: %v, %v.", resultDirection, resultDistance, table.resultDirection, table.resultDistance)
		}
	}

}

func TestGetVisibleAsteroids(t *testing.T) {
	tables := []struct {
		asteroids []asteroid
		ast       asteroid
		result    paths
	}{
		{
			[]asteroid{
				{1.5, 0.5}, {4.5, 0.5},
				{0.5, 2.5}, {1.5, 2.5}, {2.5, 2.5}, {3.5, 2.5}, {4.5, 2.5},
				{4.5, 3.5},
				{3.5, 4.5}, {4.5, 4.5},
			},
			asteroid{4.5, 4.5},
			paths{
				-2.5535900500422257: 3.605551275463989,
				-2.356194490192345:  2.8284271247461903,
				-2.0344439357957027: 4.47213595499958,
				-1.5707963267948966: 2,
				-1.3258176636680323: 4.123105625617661,
				-1.1071487177940904: 2.23606797749979,
				-0.7853981633974483: 1.4142135623730951,
				0:                   1},
		},
	}

	for _, table := range tables {
		result := getVisibleAsteroids(table.asteroids, table.ast)
		if !reflect.DeepEqual(table.result, result) {
			t.Errorf("Map was incorrect, got: %v, want: %v.", result, table.result)
		}
	}
}

// func TestGetBestAsteroid(t *testing.T) {
// 	tables := []struct {
// 		asteroids []asteroid
// 		path      paths
// 	}{
// 		{
// 			[]asteroid{
// 				{1.5, 0.5}, {4.5, 0.5},
// 				{0.5, 2.5}, {1.5, 2.5}, {2.5, 2.5}, {3.5, 2.5}, {4.5, 2.5},
// 				{4.5, 3.5},
// 				{3.5, 4.5}, {4.5, 4.5},
// 			},
// 			asteroid{3.5, 4.5},
// 			paths{
// 				-2.5535900500422257: 3.605551275463989,
// 				-2.356194490192345:  2.8284271247461903,
// 				-2.0344439357957027: 4.47213595499958,
// 				-1.5707963267948966: 2,
// 				-1.3258176636680323: 4.123105625617661,
// 				-1.1071487177940904: 2.23606797749979,
// 				-0.7853981633974483: 1.4142135623730951,
// 				0:                   1},
// 		},
// 	}

// 	for _, table := range tables {
// 		ast, p, best := getBestAsteroid(table.asteroids)
// 		if ast != table.ast || p != table.path || best != table.best {
// 			t.Errorf("Output was incorrect, got: %v, %v, %v, want: %v, %v, %v.", ast, p, best, table.ast, table.path, table.best)
// 		}
// 	}
// }
