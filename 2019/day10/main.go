package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"reflect"
	"sort"

	geo "github.com/paulmach/go.geo"
)

type asteroid struct {
	x float64
	y float64
}

type paths map[float64]struct {
	distance float64
	a        asteroid
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
	asteroids := parseMap(input)
	asteroidMap := make(map[asteroid]paths)
	for _, aster := range asteroids {
		asteroidMap[aster] = getVisibleAsteroids(asteroids, aster)
	}
	a, p, nr := getBestAsteroid(asteroidMap)
	// fmt.Println(p)

	fmt.Printf("Asteroid %v has %v in line of sight \n", a, nr)
	twoHundredth := vaporize(p, a)
	fmt.Println(twoHundredth)
	// answer := int(twoHundredthAsteroid.x)*100 + int(twoHundredthAsteroid.y)
	// fmt.Println("Answer part 2:", answer, twoHundredthAsteroid)
}

func parseMap(input []string) []asteroid {
	var asteroids []asteroid
	for i, line := range input {
		for j, element := range line {
			if string(element) == "#" {
				astY := float64(i) + 0.5
				astX := float64(j) + 0.5
				asteroids = append(asteroids, asteroid{x: astX, y: astY})
			}
		}
	}
	return asteroids
}

func vaporize(p paths, a asteroid) float64 {
	keys := make([]float64, len(p))

	i := 0
	for k := range p {
		keys[i] = k
		i++
	}
	sort.Float64s(keys)
	startDegree := -90.0
	var twohundrethAsteroid float64
	index := sort.Search(len(keys), func(i int) bool { return keys[i] >= startDegree })
	for i := 0; i < len(keys); i++ {
		if i == 199 {
			twohundrethAsteroid = keys[index]
			fmt.Println("Asteroid at end of line", keys[index], p[keys[index]])
		}
		// if reflect.DeepEqual(p[keys[index]].a, asteroid{x: 8.5, y: 2.5}) {
		// 	fmt.Println("Hit", i)
		// }
		// if index-1 < 0 {
		// 	index = len(keys) - 1
		// } else {
		index = (index + 1) % len(keys)
		// }
	}
	return twohundrethAsteroid
}

func getAsteroidFromLine(a asteroid, asteroids []asteroid, direction float64, distance float64) asteroid {
	for _, b := range asteroids {
		dir, dist := getLine(a, b)
		if dir == direction && dist == distance {
			return b
		}
	}
	return a
}

// func makeAsteroidMap(asteroids []asteroid) map[asteroid]paths {
// 	asteroidMap := make(map[asteroid]paths)
// 	for _, a := range asteroids {
// 		p := getVisibleAsteroids(asteroids, a)
// 		asteroidMap[a] = p
// 	}
// 	return asteroidMap
// }

func getBestAsteroid(asteroids map[asteroid]paths) (asteroid, paths, int) {
	var best int
	var p paths
	var a asteroid
	for ast, visibleAsteroids := range asteroids {
		if len(visibleAsteroids) > best {
			best = len(visibleAsteroids)
			a = ast
			p = visibleAsteroids
		}
	}
	return a, p, best
}

func getVisibleAsteroids(asteroids []asteroid, f asteroid) paths {
	asteroidMap := make(paths)
	for _, b := range asteroids {
		if reflect.DeepEqual(f, b) {
			continue
		}
		direction, distance := getLine(f, b)
		if val, ok := asteroidMap[direction]; ok {
			if val.distance > distance {
				p := struct {
					distance float64
					a        asteroid
				}{distance: distance, a: b}
				asteroidMap[direction] = p
			}
		} else {
			p := struct {
				distance float64
				a        asteroid
			}{distance: distance, a: b}
			asteroidMap[direction] = p
		}
	}
	return asteroidMap
}

func getLine(a asteroid, b asteroid) (float64, float64) {
	astA := geo.NewPoint(a.x, a.y)
	astB := geo.NewPoint(b.x, b.y)
	line := geo.NewLine(astA, astB)
	direction := line.Direction() * (180.0 / math.Pi)
	return direction, line.Distance()
}
