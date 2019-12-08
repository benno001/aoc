package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		layerWidth := 25
		layerHeight := 6
		layers := parseLayers(scanner.Text(), []int{layerWidth, layerHeight})
		fmt.Println(layers)
		part1(layers)
		part2(layers, []int{layerWidth, layerHeight})
	}
}

func part1(layers map[int][][]int) {
	zeros := 200
	var calc int
	for i := range layers {
		result := countNumbers(layers[i])
		if result[0] < zeros {
			zeros = result[0]
			calc = result[1] * result[2]
		}
	}
	fmt.Println("Answer part 1:", calc)
}

func part2(layers map[int][][]int, layerDimensions []int) {
	decodedImage := decodeImage(layers, layerDimensions)
	fmt.Println("Answer part 2")
	for _, row := range decodedImage {
		fmt.Println(row)
	}
}

func decodeImage(layers map[int][][]int, layerDimensions []int) [][]int {
	decodedImage := make([][]int, layerDimensions[0])
	for i := len(layers); i >= 0; i-- {
		layer := layers[i]
		decodedLayer := make([][]int, layerDimensions[0])
		for j, row := range layer {
			newRow := make([]int, len(row))
			for k, element := range row {
				if element != 2 {
					newRow[k] = element
				} else {
					if i == len(layers)-1 {
						newRow[k] = 2
					} else {
						newRow[k] = decodedImage[j][k]
					}
				}
			}
			decodedLayer[j] = newRow
		}
		decodedImage = decodedLayer
	}
	return decodedImage
}

func parseLayers(input string, layerDimensions []int) map[int][][]int {
	layers := make(map[int][][]int)
	var counter int
	for i := 0; i < len(input); i += layerDimensions[0] * layerDimensions[1] {
		layer := parseLayer(input[i:i+layerDimensions[0]*layerDimensions[1]], layerDimensions)
		layers[counter] = layer
		counter++
	}
	return layers
}

func parseLayer(input string, layerDimensions []int) [][]int {
	var layer [][]int
	for i := 0; i < len(input); i += layerDimensions[0] {
		var row []int
		for _, v := range input[i : i+layerDimensions[0]] {
			nr, _ := strconv.Atoi(string(v))
			row = append(row, nr)
		}
		layer = append(layer, row)
	}
	return layer
}

func countNumbers(layer [][]int) map[int]int {
	numbers := make(map[int]int)
	for _, i := range layer {
		for _, j := range i {
			numbers[j] = numbers[j] + 1
		}
	}
	return numbers
}
