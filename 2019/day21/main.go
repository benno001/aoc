package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const positionMode, immediateMode, relativeMode = "0", "1", "2"

// ASCIITable provides translation from ascii to code
var ASCIITable = map[rune]int{
	' ': 32,
	'#': 35,
	'.': 46,
	'@': 64,
	'A': 65,
	'B': 66,
	'C': 67,
	'D': 68,
	'E': 69,
	'F': 70,
	'G': 71,
	'H': 72,
	'I': 73,
	'J': 74,
	'K': 75,
	'L': 76,
	'M': 77,
	'N': 78,
	'O': 79,
	'P': 80,
	'Q': 81,
	'R': 82,
	'S': 83,
	'T': 84,
	'U': 85,
	'V': 86,
	'W': 87,
	'X': 88,
	'Y': 89,
	'Z': 90,
}

// invertedASCIITable provides translation from code to ASCII
var invertedASCIITable = map[int]rune{
	10:  '\n',
	32:  ' ',
	35:  '#',
	39:  '\'',
	46:  '.',
	58:  ':',
	64:  '@',
	65:  'A',
	66:  'B',
	67:  'C',
	68:  'D',
	69:  'E',
	70:  'F',
	71:  'G',
	72:  'H',
	73:  'I',
	74:  'J',
	75:  'K',
	76:  'L',
	77:  'M',
	78:  'N',
	79:  'O',
	80:  'P',
	81:  'Q',
	82:  'R',
	83:  'S',
	84:  'T',
	85:  'U',
	86:  'V',
	87:  'W',
	88:  'X',
	89:  'Y',
	90:  'Z',
	97:  'a',
	98:  'b',
	99:  'c',
	100: 'd',
	101: 'e',
	102: 'f',
	103: 'g',
	104: 'h',
	105: 'i',
	106: 'j',
	107: 'k',
	108: 'l',
	109: 'm',
	110: 'n',
	111: 'o',
	112: 'p',
	114: 'r',
	115: 's',
	116: 't',
	117: 'u',
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
	intcodes, _ := parseProgram(input)
	intc := make([]int, 10000)
	intc2 := make([]int, 10000)
	copy(intc[0:], intcodes)
	copy(intc2[0:], intcodes)
	instr := []string{"NOT A T", "NOT C J", "OR T J", "AND D J", "WALK"}
	instructions := decodeASCII(instr)
	fmt.Println(instructions)
	output := processOpcodes(intc, instructions)
	fmt.Println(decodeASCIICode(output))

	// (D && !(A && B && (C || !H)))
	instr = []string{"NOT A J", "NOT B T", "OR T J", "NOT C T", "OR T J", "AND D J", "AND E T", "OR H T", "AND T J", "RUN"}
	instructions = decodeASCII(instr)
	output = processOpcodes(intc2, instructions)
	fmt.Println(decodeASCIICode(output))
}

func decodeASCII(input []string) (output []int) {
	for _, s := range input {
		for _, char := range s {
			val, ok := ASCIITable[char]
			if !ok {
				log.Fatal("Could not encode: ", char)
			}
			output = append(output, val)
		}
		output = append(output, 10)
	}
	return output
}

func decodeASCIICode(input []int) (output []string) {
	for _, i := range input {
		val, ok := invertedASCIITable[i]
		if !ok {
			fmt.Println("Could not decode: ", i, "answer?")
		}
		output = append(output, string(val))
	}
	return output
}

func parseProgram(input []string) ([]int, error) {
	for _, s := range input {
		r := csv.NewReader(strings.NewReader(s))
		result, _ := r.ReadAll()
		for _, record := range result {
			var intcodes []int
			for _, code := range record {
				intcode, err := strconv.Atoi(code)
				intcodes = append(intcodes, intcode)
				check(err)
			}
			return intcodes, nil
		}
	}
	return []int{0}, errors.New("Error parsing input")
}

func processOpcodes(intcodes []int, input []int) (output []int) {
	relativeBase := 0
	increase := 4
	inputTracker := 0
	for position := 0; position < len(intcodes); position += increase {
		paddedIntCode := fmt.Sprintf("%05d", intcodes[position])
		instruction, _ := strconv.Atoi(paddedIntCode[3:5])

		if instruction == 99 {
			break
		}

		parameterMode := map[int]string{
			1: paddedIntCode[2:3],
			2: paddedIntCode[1:2],
			3: paddedIntCode[0:1],
		}
		getParameter := func(pos int) int {
			switch parameterMode[pos] {
			case positionMode:
				return intcodes[intcodes[position+pos]]
			case immediateMode:
				return intcodes[position+pos]
			case relativeMode:
				return intcodes[relativeBase+intcodes[position+pos]]
			}
			return 0
		}

		putParameter := func(pos int, value int) {
			switch parameterMode[pos] {
			case positionMode:
				intcodes[intcodes[position+pos]] = value
				return
			case relativeMode:
				intcodes[relativeBase+intcodes[position+pos]] = value
				return
			}
		}

		if instruction == 1 || instruction == 2 {
			if instruction == 1 {
				putParameter(3, getParameter(1)+getParameter(2))
			} else if instruction == 2 {
				putParameter(3, getParameter(1)*getParameter(2))
			}
			increase = 4
		}
		if instruction == 3 {
			putParameter(1, input[inputTracker])
			inputTracker++
			increase = 2
		}
		if instruction == 4 {
			output = append(output, getParameter(1))
			// fmt.Println("Answer: ", getParameter(1))
			increase = 2
		}
		if instruction == 5 || instruction == 6 {
			var newPosition int
			if (instruction == 5 && getParameter(1) != 0) || (instruction == 6 && getParameter(1) == 0) {
				newPosition = getParameter(2)
			} else {
				newPosition = position + 3
			}
			increase = newPosition - position
		}
		if instruction == 7 || instruction == 8 {
			if (instruction == 7 && getParameter(1) < getParameter(2)) || (instruction == 8 && getParameter(1) == getParameter(2)) {
				putParameter(3, 1)
			} else {
				putParameter(3, 0)
			}
			increase = 4
		}
		if instruction == 9 {
			relativeBase = relativeBase + getParameter(1)
			increase = 2
		}
	}
	return output
}
