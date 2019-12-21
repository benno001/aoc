package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"gopkg.in/karalabe/cookiejar.v2/graph"
	"gopkg.in/karalabe/cookiejar.v2/graph/bfs"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
	grid := parseInput(input)
	// visualizeGrid(grid)
	g, innerPortals, outerPortals := createGraph(grid)
	connectPortals(g, innerPortals, outerPortals)
	fmt.Println(innerPortals, outerPortals)
	path, steps := getShortestPathBfs(g, outerPortals["AA"], outerPortals["ZZ"])
	fmt.Println(path, steps)
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}

func invertMap(portals map[string]int) map[int]string {
	p := make(map[int]string)
	for k, v := range portals {
		p[v] = k
	}
	return p
}

func getShortestPathBfs(g *graph.Graph, sourceVertice int, destinationVertice int) ([]int, int) {
	b := bfs.New(g, sourceVertice)
	path := b.Path(destinationVertice)
	return path, len(path) - 1
}

func visualizeGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

// Parse input
func parseInput(input []string) (grid [][]string) {
	grid = make([][]string, len(input))
	for i := range grid {
		grid[i] = make([]string, len(input[0]))
	}
	for i, row := range input {
		for j, char := range row {
			c := string(char)
			grid[i][j] = c
		}
	}
	return grid
}

func connectPortals(g *graph.Graph, innerPortals map[string]int, outerPortals map[string]int) {
	for portal, v1 := range innerPortals {
		if v2, ok := outerPortals[portal]; ok {
			g.Connect(v1, v2)
		}
	}
}

func createGraph(grid [][]string) (g *graph.Graph, innerPortals map[string]int, outerPortals map[string]int) {
	innerPortals = make(map[string]int)
	outerPortals = make(map[string]int)
	verticeAmount := len(grid) * len(grid) * 10
	size := len(grid[0])
	g = graph.New(verticeAmount)
	for i, row := range grid {
		isInner := false
		for j, cell := range row {
			if cell == "#" || cell == "." {
				isInner = true
			}
			if cell != "#" && cell != " " {
				if cell != "." {
					pos := i*size + j
					if j >= len(row) - 3 {
						isInner = false
					}
					if j == len(row)-1 || i == len(grid)-1 {
						continue
					}
					portal, _ := regexp.MatchString("[A-Z]", cell)
					
					right := grid[i][j+1]
					portalRight, _ := regexp.MatchString("[A-Z]", right)

					down := grid[i+1][j]
					portalDown, _ := regexp.MatchString("[A-Z]", down)

					if portal && portalRight {
						portalString := strings.Join([]string{cell, right}, "")
						pos2 := i*size + (j+1)
						pos3 := 0
						if j >= 2 {
							left := grid[i][j-1]
							left2 := grid[i][j-2]
							if left == "." {
								pos3 = i*size +(j-1)
							} else if left2 == "." && left != "." {
								pos3 = i*size +(j-2)
							}
						}
						if j < len(row) -2 {
							right2 := grid[i][j+2]
							if right == "."{
								pos3 = i*size +(j+1)
							} else if right2 == "." && right != "." {
								pos3 = i*size +(j+2)
							}
						}
						g.Connect(pos, pos2)
						if isInner {
							innerPortals[portalString] = pos3
						} else {
							outerPortals[portalString] = pos3
						}
					} else if portal && portalDown {
						portalString := strings.Join([]string{cell, down}, "")
						pos2 := (i+1)*size +j
						pos3 := 0
						if i < len(grid) -2 {
							down2 := grid[i+2][j]
							if down == "." {
								pos3 = (i+1)*size +j
							} else if down2 == "." && down != "." {
								pos3 = (i+2)*size +j
							}
						}
						if i >= 2 {
							up := grid[i-1][j]
							up2 := grid[i-2][j]
							if up == "." {
								pos3 = (i-1)*size +j
							} else if up2 == "." && up != "." {
								pos3 = (i-2)*size +j
							}
						}
						g.Connect(pos, pos2)
						if isInner {
							innerPortals[portalString] = pos3
						} else {
							outerPortals[portalString] = pos3
						}
					}
				}
				if i-1 > 0 {
					up := grid[i-1][j]
					portal, _ := regexp.MatchString("[A-Z]", up)
					if up == "." || portal {
						g.Connect(i*size+j, (i-1)*size+j)
					}
				}
				if i+1 < len(grid) {
					down := grid[i+1][j]
					portal, _ := regexp.MatchString("[A-Z]", down)
					if down == "." || portal {
						g.Connect(i*size+j, (i+1)*size+j)
					}
				}
				if j-1 > 0 {
					left := grid[i][j-1]
					portal, _ := regexp.MatchString("[A-Z]", left)
					if left == "." || portal {
						g.Connect(i*size+j, i*size+(j-1))
					}
				}
				if j+1 < len(row) {
					right := grid[i][j+1]
					portal, _ := regexp.MatchString("[A-Z]", right)
					if right == "." || portal {
						g.Connect(i*size+j, i*size+(j+1))
					}
				}
			}
		}
	}
	// fmt.Println(keys)
	return g, innerPortals, outerPortals
}
