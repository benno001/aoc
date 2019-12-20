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
	grid := parseInput(input)
	g, keys, doors := createGraph(grid)
	visualizeGrid(grid)
	precomp := precomputeDistances(g, keys)
	minSteps := distanceToCollectKeys(g, "@", keys, make(map[[32]byte]int), doors, precomp)
	fmt.Println("Answer part 1:", minSteps)
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
			if cell != "#" {
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
	return distanceToCollectKeys(g, "@", keys, make(map[[32]byte]int), doors, precomputedDistances)
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

func distanceToCollectKeys(g *graph.Graph, currentKey string, newKeys map[string]int, cache map[[32]byte]int, newDoors map[string]int, precomputedDistances map[[32]byte]precomputedDistance) int {
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

	for key := range k {
		delete(doors, strings.ToUpper(key))
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

		fmt.Println("dist for", currentKey, key, precomp.dist, k)
		// delete(nKeys, key)
		distance := precomp.dist + distanceToCollectKeys(g, key, nKeys, cache, nDoors, precomputedDistances)
		if distance < minSteps {
			minSteps = distance
			// fmt.Println(minSteps)
		}
	}
	cache[sha256sum] = minSteps
	return minSteps
}
