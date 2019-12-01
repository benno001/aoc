package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	total := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(strings.ToUpper(scanner.Text()))
		total += calculateFuel(mass, 0)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
	fmt.Printf("Total is %d", total)
}

func calculateFuel(mass int, required int) int {
	extraFuel := mass / 3 - 2
	if extraFuel <= 0 {
		return required
	}
	return calculateFuel(extraFuel, required + extraFuel)
}
