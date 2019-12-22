package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"

	"gopkg.in/karalabe/cookiejar.v2/graph"
	"gopkg.in/karalabe/cookiejar.v2/graph/bfs"
)

type precomputedDistance struct {
	dist int
	kip  map[string]int
}

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
	// grid := parseInput(input)
	// g, keys, doors := createGraph(grid)
	// visualizeGrid(grid)
	// precomp := precomputeDistances(g, keys)
	// minSteps := distanceToCollectKeys(g, "@", keys, make(map[[32]byte]int), doors, precomp, 0)
	// fmt.Println("Answer part 1:", minSteps)

	q1, q2, q3, q4 := parseQuadrants(input)
	visualizeGrid(q1)
	visualizeGrid(q2)
	visualizeGrid(q3)
	visualizeGrid(q4)
	g1, k1, d1 := createGraph(q1)
	removeIrrelevantDoors(k1, d1)
	precomp1 := precomputeDistances(g1, k1)
	minSteps1 := distanceToCollectKeys(g1, "@", k1, make(map[[32]byte]int), d1, precomp1, 0)

	g2, k2, d2 := createGraph(q2)
	removeIrrelevantDoors(k2, d2)
	precomp2 := precomputeDistances(g2, k2)
	minSteps2 := distanceToCollectKeys(g2, "@", k2, make(map[[32]byte]int), d2, precomp2, 0)

	g3, k3, d3 := createGraph(q3)
	removeIrrelevantDoors(k3, d3)
	precomp3 := precomputeDistances(g3, k3)
	minSteps3 := distanceToCollectKeys(g3, "@", k3, make(map[[32]byte]int), d3, precomp3, 0)

	g4, k4, d4 := createGraph(q4)
	removeIrrelevantDoors(k4, d4)
	precomp4 := precomputeDistances(g4, k4)
	minSteps4 := distanceToCollectKeys(g4, "@", k4, make(map[[32]byte]int), d4, precomp4, 0)

	fmt.Println("Answer part 2:", minSteps1+minSteps2+minSteps3+minSteps4)

}

func removeIrrelevantDoors(keys map[string]int, doors map[string]int) {
	for d := range doors {
		if _, ok := keys[strings.ToUpper(d)]; !ok {
			delete(doors, d)
		}
	}
}

func parseInput(input []string) (grid [][]string) {
	grid = make([][]string, len(input))
	for i := range grid {
		grid[i] = make([]string, len(input))
	}
	for i, row := range input {
		for j, char := range row {
			c := string(char)
			grid[i][j] = c
		}
	}
	return grid
}

func parseQuadrants(input []string) (quadrant1 [][]string, quadrant2 [][]string, quadrant3 [][]string, quadrant4 [][]string) {
	quadrant1 = make([][]string, len(input))
	quadrant2 = make([][]string, len(input))
	quadrant3 = make([][]string, len(input))
	quadrant4 = make([][]string, len(input))
	for i := range quadrant1 {
		quadrant1[i] = make([]string, len(input))
		quadrant2[i] = make([]string, len(input))
		quadrant3[i] = make([]string, len(input))
		quadrant4[i] = make([]string, len(input))
	}
	for i, row := range input {
		for j, char := range row {
			c := string(char)
			if i < len(input)/2 && j < len(input[i])/2 {
				quadrant1[i][j] = c
			}
			if i < len(input)/2 && j >= len(input[i])/2 {
				quadrant2[i][j] = c
			}
			if i >= len(input)/2 && j < len(input[i])/2 {
				quadrant3[i][j] = c
			}
			if i >= len(input)/2 && j >= len(input[i])/2 {
				quadrant4[i][j] = c
			}
		}
	}
	return quadrant1, quadrant2, quadrant3, quadrant4
}

func visualizeGrid(grid [][]string) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

// Iterate over all vertices, get them connected.
func createGraph(grid [][]string) (g *graph.Graph, keys map[string]int, doors map[string]int) {
	keys = make(map[string]int)
	doors = make(map[string]int)
	verticeAmount := len(grid) * len(grid) * 10
	size := len(grid[0])
	g = graph.New(verticeAmount)
	for i, row := range grid {
		for j, cell := range row {
			if cell != "#" && cell != " " && cell != "" {
				if cell != "." {
					pos := i*size + j
					// If the cell is an item, also add it to the map
					door, _ := regexp.MatchString("[A-Z]", cell)
					key, _ := regexp.MatchString("[@a-z]", cell)
					if door {
						doors[cell] = pos
					} else if key {
						keys[cell] = pos
					}
				}
				up := grid[i-1][j]
				if up != "#" {
					g.Connect(i*size+j, (i-1)*size+j)
				}

				down := grid[i+1][j]

				if down != "#" {
					g.Connect(i*size+j, (i+1)*size+j)
				}

				left := grid[i][j-1]
				if left != "#" {
					g.Connect(i*size+j, i*size+(j-1))
				}

				right := grid[i][j+1]
				if right != "#" {
					g.Connect(i*size+j, i*size+(j+1))
				}
			}
		}
	}
	// fmt.Println(keys)
	return g, keys, doors
}

func traverseMaze(g *graph.Graph, doors map[string]int, keys map[string]int, precomputedDistances map[[32]byte]precomputedDistance) int {
	return distanceToCollectKeys(g, "@", keys, make(map[[32]byte]int), doors, precomputedDistances, 0)
}

func getPossibleKeys(g *graph.Graph, player int, keys map[string]int, doors map[string]int) map[string]int {
	possibleKeys := make(map[string]int)
	for key, vertice := range keys {
		if !keyIsBlocked(g, player, vertice, doors) {
			// fmt.Println("k,v", key, vertice, doors)
			b := bfs.New(g, player)
			path := b.Path(vertice)
			possibleKeys[key] = len(path) - 1
		}
	}
	return possibleKeys
}

func keyIsBlocked(g *graph.Graph, player int, key int, doors map[string]int) bool {
	b := bfs.New(g, player)
	path := b.Path(key)
	for _, door := range doors {
		for _, vertice := range path {
			if door == vertice {
				return true
			}
		}
	}
	return false
}

func precomputeDistances(g *graph.Graph, keys map[string]int) map[[32]byte]precomputedDistance {
	precomputedDistances := make(map[[32]byte]precomputedDistance)
	for a, v1 := range keys {
		for b, v2 := range keys {
			if a != b {
				search := bfs.New(g, v1)
				path := search.Path(v2)
				kip := keysInPath(keys, path)
				shaMaterial := append([]byte(a), []byte(b)...)
				sha256 := sha256.Sum256(shaMaterial)
				precomputedDistances[sha256] = precomputedDistance{dist: len(path) - 1, kip: kip}
			}
		}
	}
	return precomputedDistances
}

func distanceToKey(g *graph.Graph, a int, b string, keys map[string]int) (int, []int) {
	verticeB := keys[b]
	// fmt.Println(keys, a, b, verticeB)
	search := bfs.New(g, a)
	path := search.Path(verticeB)
	// fmt.Println(path)
	if len(path) == 0 {
		log.Fatal("Encountered erroneous path")
	}
	return len(path) - 1, path
}

func keysInPath(keys map[string]int, path []int) map[string]int {
	keysInPath := make(map[string]int)
	for k, v := range keys {
		for i, vertice := range path {
			if v == vertice && i != len(path)-1 {
				keysInPath[k] = v
			}
		}
	}
	return keysInPath
}

func distanceToCollectKeys(g *graph.Graph, currentKey string, newKeys map[string]int, cache map[[32]byte]int, newDoors map[string]int, precomputedDistances map[[32]byte]precomputedDistance, depth int) int {
	keys := make(map[string]int)
	for k, v := range newKeys {
		keys[k] = v
	}
	doors := make(map[string]int)
	for k, v := range newDoors {
		doors[k] = v
	}
	// cache := make(map[string]cachedKey)
	// for k, v := range newCache {
	// 	cache[k] = cachedKey{keys: v.keys, dist: v.dist}
	// }
	if len(keys) == 0 {
		return 0
	}
	currentVertice := keys[currentKey]
	delete(keys, currentKey)
	delete(doors, strings.ToUpper(currentKey))
	keyList := make([]string, len(keys))
	i := 0
	for k := range keys {
		keyList[i] = k
		i++
	}
	sort.Strings(keyList)
	keysString := strings.Join(keyList, "")
	shaMaterial := append([]byte(currentKey), []byte(keysString)...)
	sha256sum := sha256.Sum256(shaMaterial)
	if _, ok := cache[sha256sum]; ok {
		return cache[sha256sum]
	}

	minSteps := 9999999999
	k := getPossibleKeys(g, currentVertice, keys, doors)
	if len(k) == 0 {
		return 0
	}

	j := 0
	newKeyList := make([]string, len(k))
	for newKey := range k {
		newKeyList[j] = newKey
		j++
	}
	sort.Strings(newKeyList)
	// fmt.Println(newKeyList)

	for _, key := range newKeyList {
		// delete(doors, strings.ToUpper(key))
		nKeys := make(map[string]int)
		for a, v := range keys {
			nKeys[a] = v
		}
		nDoors := make(map[string]int)
		for b, v := range doors {
			nDoors[b] = v
		}

		keyMaterial := append([]byte(currentKey), []byte(key)...)
		keySha256 := sha256.Sum256(keyMaterial)

		precomp := precomputedDistances[keySha256]

		for kip := range precomp.kip {
			delete(nKeys, kip)
			delete(nDoors, strings.ToUpper(kip))
		}

		fmt.Println(strings.Repeat("  ", depth), currentKey, key, precomp.dist, k, nDoors)
		// delete(nKeys, key)
		distance := precomp.dist + distanceToCollectKeys(g, key, nKeys, cache, nDoors, precomputedDistances, depth+1)
		if distance < minSteps {
			minSteps = distance
			// fmt.Println(minSteps)
		}
	}
	cache[sha256sum] = minSteps
	return minSteps
}
