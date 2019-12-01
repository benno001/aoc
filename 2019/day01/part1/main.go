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
		total += calculateFuel(mass)
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
	}
	fmt.Printf("Total is %d", total)
}

func calculateFuel(mass int) int {
	return mass/3 - 2
}
