package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"io/ioutil"
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
	'\n': 10,
	' ':  32,
	'!':  33,
	'#':  35,
	'.':  46,
	'?':  63,
	'@':  64,
	'A':  65,
	'B':  66,
	'C':  67,
	'D':  68,
	'E':  69,
	'F':  70,
	'G':  71,
	'H':  72,
	'I':  73,
	'J':  74,
	'K':  75,
	'L':  76,
	'M':  77,
	'N':  78,
	'O':  79,
	'P':  80,
	'Q':  81,
	'R':  82,
	'S':  83,
	'T':  84,
	'U':  85,
	'V':  86,
	'W':  87,
	'X':  88,
	'Y':  89,
	'Z':  90,
	'a':  97,
	'b':  98,
	'c':  99,
	'd':  100,
	'e':  101,
	'f':  102,
	'g':  103,
	'h':  104,
	'i':  105,
	'j':  106,
	'k':  107,
	'l':  108,
	'm':  109,
	'n':  110,
	'o':  111,
	'p':  112,
	'r':  114,
	's':  115,
	't':  116,
	'u':  117,
	'v':  118,
	'w':  119,
	'x':  120,
	'y':  121,
	'z':  122,
}

// invertedASCIITable provides translation from code to ASCII
var invertedASCIITable = map[int]rune{
	10:  '\n',
	32:  ' ',
	33:  '!',
	34:  '"',
	35:  '#',
	39:  '\'',
	44:  ',',
	45:  '-',
	46:  '.',
	50:  '2',
	51:  '3',
	54:  '6',
	55:  '7',
	58:  ':',
	59:  ';',
	61:  '=',
	63:  '?',
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
	113: 'q',
	114: 'r',
	115: 's',
	116: 't',
	117: 'u',
	118: 'v',
	119: 'w',
	120: 'x',
	121: 'y',
	122: 'z',
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	data, err := ioutil.ReadFile("input")
	check(err)
	intcodes, _ := parseProgram(string(data))
	intc := make([]int, 10000)
	intc2 := make([]int, 10000)
	copy(intc[0:], intcodes)
	copy(intc2[0:], intcodes)
	// input := 1
	processOpcodes(intc)
	// input = 2
	// processOpcodes(intc2, input)
}

func parseProgram(input string) ([]int, error) {
	r := csv.NewReader(strings.NewReader(input))
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
	return []int{0}, errors.New("Error parsing input")
}

func decodeASCII(input string) (output []int) {
	for _, char := range input {
		val, ok := ASCIITable[char]
		if !ok {
			log.Fatal("Could not encode: ", char)
		}
		output = append(output, val)
	}
	// output = append(output, 10)
	return output
}

func decodeASCIICode(input []int) (output string) {
	var out []string
	for n, i := range input {
		if (n == 0 && i == 10) || n == len(input)-1 {
			continue
		}
		val, ok := invertedASCIITable[i]
		if !ok {
			fmt.Println("Could not decode: ", i, "answer?")
		}
		out = append(out, string(val))
	}
	return strings.Join(out, "")
}

func askInput() []int {
	reader := bufio.NewReader(os.Stdin)
	// reader := bufio.NewReader(os.Stdin)
	var text string
	fmt.Print(">")
	text, _ = reader.ReadString('\n')
	// text, _ :=
	return decodeASCII(text)
}

func processOpcodes(intcodes []int) {
	relativeBase := 0
	increase := 4
	var input []int
	var output []int
	inputTracker := 0
	for position := 0; position < len(intcodes); position += increase {
		paddedIntCode := fmt.Sprintf("%05d", intcodes[position])
		instruction, _ := strconv.Atoi(paddedIntCode[3:5])

		if instruction == 99 {
			fmt.Println(decodeASCIICode(output))
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
			if len(input) == 1 || inputTracker == len(input) || len(input) == 0 {
				input = askInput()
				inputTracker = 0
			}
			putParameter(1, input[inputTracker])
			inputTracker++
			increase = 2
		}
		if instruction == 4 {
			out := getParameter(1)
			output = append(output, out)
			if out == 10 && len(output) > 1 {
				if output[len(output)-2] == 10 {
					fmt.Print(decodeASCIICode(output))
					output = []int{}
				}
			} else if len(output) == 1 {
				// fmt.Println("Possible answer:", getParameter(1))
			}
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
			// fmt.Println("comparing", getParameter(1), "and", getParameter(2))
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
}
