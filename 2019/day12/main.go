package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type moon struct {
	position point
	velocity point
}

type oneAxisMoon struct {
	position int
	velocity int
}

type point struct {
	x int
	y int
	z int
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
	pristineMoons := parseInput(input)
	// moons := copy(moons, pristineMoons)
	moons := make([]moon, len(pristineMoons))
	copy(moons, pristineMoons)

	processSteps(moons, 1000)
	fmt.Println("Answer part 1:", calculateSystemEnergy(moons))

	copy(moons, pristineMoons)
	processStepsIndependent(moons)
}

func parseInput(input []string) (moons []moon) {
	for _, mString := range input {
		var m moon
		x, err := strconv.Atoi(string(mString[strings.Index(mString, "x")+2 : strings.Index(mString, ",")]))
		if err != nil {
			log.Fatal("Error converting string", err)
		}
		y, err := strconv.Atoi(string(mString[strings.Index(mString, "y")+2 : strings.LastIndex(mString, ",")]))
		if err != nil {
			log.Fatal("Error converting string", err)
		}
		z, err := strconv.Atoi(string(mString[strings.Index(mString, "z")+2 : strings.Index(mString, ">")]))
		if err != nil {
			log.Fatal("Error converting string", err)
		}
		m.position.x = x
		m.position.y = y
		m.position.z = z
		moons = append(moons, m)
	}
	return moons
}

func processSteps(moons []moon, steps int) {
	for i := 0; i < steps; i++ {
		processGravity(moons)
		processVelocity(moons)
	}
}

func processStepsIndependent(moons []moon) {
	var pristineMoonsX []oneAxisMoon
	var pristineMoonsY []oneAxisMoon
	var pristineMoonsZ []oneAxisMoon
	var moonsX []oneAxisMoon
	var moonsY []oneAxisMoon
	var moonsZ []oneAxisMoon
	for _, m := range moons {
		pristineMoonsX = append(pristineMoonsX, oneAxisMoon{position: m.position.x, velocity: 0})
		pristineMoonsY = append(pristineMoonsY, oneAxisMoon{position: m.position.y, velocity: 0})
		pristineMoonsZ = append(pristineMoonsZ, oneAxisMoon{position: m.position.z, velocity: 0})

		moonsX = append(moonsX, oneAxisMoon{position: m.position.x, velocity: 0})
		moonsY = append(moonsY, oneAxisMoon{position: m.position.y, velocity: 0})
		moonsZ = append(moonsZ, oneAxisMoon{position: m.position.z, velocity: 0})
	}
	fmt.Println("Moons: ", len(pristineMoonsX))
	metaMoons := [][]oneAxisMoon{moonsX, moonsY, moonsZ}
	metaPristineMoons := [][]oneAxisMoon{pristineMoonsX, pristineMoonsY, pristineMoonsZ}
	var outputs []int
	for i := 0; i < len(metaPristineMoons); i++ {
		out := make(chan int)
		go func(mo []oneAxisMoon, pristineMo []oneAxisMoon, o chan int) {
			for steps := 1; true; steps++ {
				processGravityIndepdendent(mo)
				processVelocityIndependent(mo)
				if reflect.DeepEqual(mo, pristineMo) {
					o <- steps
					break
				}
			}
		}(metaMoons[i], metaPristineMoons[i], out)
		output := <-out
		outputs = append(outputs, output)
	}
	fmt.Println(outputs)
	answer := LCM(outputs[0], outputs[1], outputs[2])
	fmt.Println("Answer part 2:", answer)
}

func processGravityIndepdendent(moons []oneAxisMoon) {
	for i, m := range moons {
		for _, m2 := range moons {
			if reflect.DeepEqual(m, m2) {
				continue
			}
			if m2.position > m.position {
				m.velocity++
			} else if m2.position < m.position {
				m.velocity--
			}
			moons[i] = m
		}
	}
}

// To apply gravity, consider every pair of moons.
// On each axis (x, y, and z), the velocity of each moon changes by exactly +1 or -1 to pull the moons together.
func processGravity(moons []moon) {
	for i, m := range moons {
		for _, m2 := range moons {
			if reflect.DeepEqual(m, m2) {
				continue
			}
			if m2.position.x > m.position.x {
				m.velocity.x++
			} else if m2.position.x < m.position.x {
				m.velocity.x--
			}
			if m2.position.y > m.position.y {
				m.velocity.y++
			} else if m2.position.y < m.position.y {
				m.velocity.y--
			}
			if m2.position.z > m.position.z {
				m.velocity.z++
			} else if m2.position.z < m.position.z {
				m.velocity.z--
			}
			moons[i] = m
		}
	}
}

func processVelocityIndependent(moons []oneAxisMoon) {
	for i, m := range moons {
		m.position += m.velocity
		moons[i] = m
	}
}

// Simply add the velocity of each moon to its own position
func processVelocity(moons []moon) {
	for i, m := range moons {
		m.position.x += m.velocity.x
		m.position.y += m.velocity.y
		m.position.z += m.velocity.z
		moons[i] = m
	}
}

// A moon's potential energy is the sum of the absolute values of its x, y, and z position coordinates
func calculatePotentialEnergy(m moon) int {
	x := math.Abs(float64(m.position.x))
	y := math.Abs(float64(m.position.y))
	z := math.Abs(float64(m.position.z))
	return int(x + y + z)
}

// A moon's kinetic energy is the sum of the absolute values of its velocity coordinates
func calculateKineticEnergy(m moon) int {
	x := math.Abs(float64(m.velocity.x))
	y := math.Abs(float64(m.velocity.y))
	z := math.Abs(float64(m.velocity.z))
	return int(x + y + z)
}

func calculateTotalEnergy(m moon) int {
	return calculatePotentialEnergy(m) * calculateKineticEnergy(m)
}

func calculateSystemEnergy(moons []moon) int {
	totalEnergy := 0
	for _, m := range moons {
		totalEnergy += calculateTotalEnergy(m)
	}
	return totalEnergy
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
