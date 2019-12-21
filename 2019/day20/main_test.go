package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateGraph(t *testing.T) {
	tables := []struct {
		input        []string
		innerPortals map[string]int
		outerPortals map[string]int
	}{
		{
			[]string{
				"         A           ",
				"         A           ",
				"  #######.#########  ",
				"  #######.........#  ",
				"  #######.#######.#  ",
				"  #######.#######.#  ",
				"  #######.#######.#  ",
				"  #####  B    ###.#  ",
				"BC...##  C    ###.#  ",
				"  ##.##       ###.#  ",
				"  ##...DE  F  ###.#  ",
				"  #####    G  ###.#  ",
				"  #########.#####.#  ",
				"DE..#######...###.#  ",
				"  #.#########.###.#  ",
				"FG..#########.....#  ",
				"  ###########.#####  ",
				"             Z       ",
				"             Z       ",
			},
			map[string]int{"BC": 156, "DE": 217, "FG": 221},
			map[string]int{"AA": 9, "BC": 168, "DE": 273, "FG": 315, "ZZ": 370},
		},
	}
	for _, table := range tables {
		grid := parseInput(table.input)
		_, innerPortals, outerPortals := createGraph(grid)
		if !reflect.DeepEqual(table.innerPortals, innerPortals) || !reflect.DeepEqual(table.outerPortals, outerPortals) {
			t.Errorf("Output incorrect: got %v, %v, want %v, %v.", innerPortals, outerPortals, table.innerPortals, table.outerPortals)
		}
	}
}
func TestGetPath(t *testing.T) {
	tables := []struct {
		input      []string
		pathLength int
	}{
		{
			[]string{
				"         A           ",
				"         A           ",
				"  #######.#########  ",
				"  #######.........#  ",
				"  #######.#######.#  ",
				"  #######.#######.#  ",
				"  #######.#######.#  ",
				"  #####  B    ###.#  ",
				"BC...##  C    ###.#  ",
				"  ##.##       ###.#  ",
				"  ##...DE  F  ###.#  ",
				"  #####    G  ###.#  ",
				"  #########.#####.#  ",
				"DE..#######...###.#  ",
				"  #.#########.###.#  ",
				"FG..#########.....#  ",
				"  ###########.#####  ",
				"             Z       ",
				"             Z       ",
			}, 23,
		},
	}
	for _, table := range tables {
		grid := parseInput(table.input)
		g, innerPortals, outerPortals := createGraph(grid)
		connectPortals(g, innerPortals, outerPortals)
		path, steps := getShortestPathBfs(g, outerPortals["AA"], outerPortals["ZZ"])
		if steps != table.pathLength {
			fmt.Println(path, innerPortals, outerPortals)
			t.Errorf("Output incorrect: got %v, want %v.", steps, table.pathLength)
		}
	}
}

func TestPathRecursive(t *testing.T) {
	tables := []struct {
		input []string
		steps int
	}{
		{
			[]string{
				"           Z L X W       C                   ",
				"           Z P Q B       K                   ",
				"  ###########.#.#.#.#######.###############  ",
				"  #...#.......#.#.......#.#.......#.#.#...#  ",
				"  ###.#.#.#.#.#.#.#.###.#.#.#######.#.#.###  ",
				"  #.#...#.#.#...#.#.#...#...#...#.#.......#  ",
				"  #.###.#######.###.###.#.###.###.#.#######  ",
				"  #...#.......#.#...#...#.............#...#  ",
				"  #.#########.#######.#.#######.#######.###  ",
				"  #...#.#    F       R I       Z    #.#.#.#  ",
				"  #.###.#    D       E C       H    #.#.#.#  ",
				"  #.#...#                           #...#.#  ",
				"  #.###.#                           #.###.#  ",
				"  #.#....OA                       WB..#.#..ZH",
				"  #.###.#                           #.#.#.#  ",
				"CJ......#                           #.....#  ",
				"  #######                           #######  ",
				"  #.#....CK                         #......IC",
				"  #.###.#                           #.###.#  ",
				"  #.....#                           #...#.#  ",
				"  ###.###                           #.#.#.#  ",
				"XF....#.#                         RF..#.#.#  ",
				"  #####.#                           #######  ",
				"  #......CJ                       NM..#...#  ",
				"  ###.#.#                           #.###.#  ",
				"RE....#.#                           #......RF",
				"  ###.###        X   X       L      #.#.#.#  ",
				"  #.....#        F   Q       P      #.#.#.#  ",
				"  ###.###########.###.#######.#########.###  ",
				"  #.....#...#.....#.......#...#.....#.#...#  ",
				"  #####.#.###.#######.#######.###.###.#.#.#  ",
				"  #.......#.......#.#.#.#.#...#...#...#.#.#  ",
				"  #####.###.#####.#.#.#.#.###.###.#.###.###  ",
				"  #.......#.....#.#...#...............#...#  ",
				"  #############.#.#.###.###################  ",
				"             A O F   N                       ",
				"             A A D   M                       ",
			}, 396,
		},
	}
	for _, table := range tables {
		_, steps := getPathRecursive()
		if steps != table.steps {
			t.Errorf("Output incorrect: got %v, want %v.", steps, table.steps)
		}
	}
}
