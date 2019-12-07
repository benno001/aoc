package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input []string
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	orbits := createOrbits(input)
	numOrbits := countOrbits(orbits)
	fmt.Println("Answer 1:", numOrbits)
	
	santa := getOrbitPathToRoot(orbits, "SAN")
	fmt.Println("Answer 2:", getRequiredTransfers(orbits, santa))
}

func createOrbits(input []string) map[string]string {
	orbits := make(map[string]string)
	for _, orbit := range input {
		mass := orbit[:3]
		moon := orbit[4:]
		orbits[moon] = mass
	}
	return orbits
}

func countOrbits(orbits map[string]string) int {
	var numOrbits int

	for orbit := range orbits {
		var countOrbitsOfMass int
		currentOrbit := orbit
		for {
			if currentOrbit == "COM" {
				break
			}
			countOrbitsOfMass++
			currentOrbit = orbits[currentOrbit]
		}
		numOrbits += countOrbitsOfMass
	}
	return numOrbits
}

func getOrbitPathToRoot(orbits map[string]string, mass string) map[string]int {
	orbitPath := make(map[string]int)
	currentOrbit := orbits[mass]
	for steps := 0; currentOrbit != "COM"; steps++ {
		orbitPath[currentOrbit] = steps
		currentOrbit = orbits[currentOrbit]
	}
	return orbitPath
}

func getRequiredTransfers(orbits map[string]string, orbitsOther map[string]int) int{
	var countOrbits int
	currentOrbit := orbits["YOU"]
	sameOrbit := false
	for steps := 0; !sameOrbit; steps++ {
		var orbitsFromOther int
		orbitsFromOther, sameOrbit = orbitsOther[currentOrbit]
		countOrbits = steps + orbitsFromOther
		currentOrbit = orbits[currentOrbit]
	}
	return countOrbits
}