package main

import (
	"bufio"
	"encoding/csv"
	"errors"
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
	for scanner.Scan() {
		
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			os.Exit(1)
		}
		intcodes, _ := parseProgram(scanner.Text())
		part1(intcodes)
		part2(intcodes)
	}
}

func part1(intcodes []int){
	comb := getCombinations([]int{0,1,2,3,4})
	var max int
	for _, c := range comb {
		res := runAmplifierSetup(intcodes, c)
		if res > max {
			max = res
		}
	}
	fmt.Println("Answer part 1:", max)
}

func part2(intcodes []int) {
	comb := getCombinations([]int{5,6,7,8,9})
	max := 0
	for _, c := range comb {
		res := runFeedBackAmplifierSetup(intcodes, c)
		if res > max {
			max = res
		}
	}
	fmt.Println("Answer part 2:", max)
}

func runAmplifierSetup(intcodes []int, phases []int) int {
	amplifiers := []string{"A","B","C","D","E"}
	var thrusterResult int
	var err error
	for i, amp := range amplifiers {
		phase := phases[i]
		if amp == "A" {
			thrusterResult = 0
		}
		thrusterResult, err = processOpcodes(intcodes, phase, thrusterResult)
		if err != nil {
			break
		}
	}
	return thrusterResult
}

func runFeedBackAmplifierSetup(intcodes []int, phases []int) int {
	in := make(chan int, 1)
	in <- 0
	out := in
	for _, phase := range phases {
		out = runProgram(intcodes, phase, out)
	}
	var output int
	for output = range out {
		in <- output
	}
	return output
}

func runProgram(intcodes []int, phase int, signals <-chan int) chan int {
	intcodes = copiedInstructions(intcodes)
	in := make(chan int)
	out := make(chan int)
	go func() {
		in <- phase
		for signal := range signals {
			in <- signal
		}
		close(in)
	}()
	go processOpcodesFeedBack(intcodes, in, out)
	return out
}

func copiedInstructions(intcodes []int) []int {
	intcodesCopy := make([]int, len(intcodes))
	copy(intcodesCopy, intcodes)
	return intcodesCopy
}

func getCombinations(arr []int) [][]int {
    var helper func([]int, int)
    res := [][]int{}

    helper = func(arr []int, n int) {
        if n == 1 {
            tmp := make([]int, len(arr))
            copy(tmp, arr)
            res = append(res, tmp)
        } else {
            for i := 0; i < n; i++ {
                helper(arr, n-1)
                if n%2 == 1 {
                    tmp := arr[i]
                    arr[i] = arr[n-1]
                    arr[n-1] = tmp
                } else {
                    tmp := arr[0]
                    arr[0] = arr[n-1]
                    arr[n-1] = tmp
                }
            }
        }
    }
    helper(arr, len(arr))
    return res
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

func processOpcodes(intcodes []int, phase int, input int) (int, error) {
	increase := 4
	firstInput := true
	instructionMap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
	}
	for position := 0; position < len(intcodes); position += increase {
		instr := intcodes[position]
		instruction, immediateOne, immediateTwo := processInstruction(instr)
		if _, ok := instructionMap[instruction]; !ok || instruction == 99 {
			return 0, errors.New("halted")
			break
		}
		var numberOne int
		var numberTwo int
		if instruction == 1 || instruction == 2 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			out := intcodes[position+3]
			output, err := processOneTwoInstruction(instruction, numberOne, numberTwo)
			check(err)
			intcodes[out] = output
			increase = 4
		}
		if instruction == 3 {
			numberOne = intcodes[position+1]
			if firstInput {
				processOpcodeThree(intcodes, numberOne, phase)
				firstInput = false
				} else {
				processOpcodeThree(intcodes, numberOne, input)
			}
			increase = 2
		}
		if instruction == 4 {
			numberOne = intcodes[position+1]
			res := processOpcodeFour(intcodes, numberOne)
			// fmt.Println("Answer: ", res)
			increase = 2
			return res, nil
		}
		if instruction == 5 || instruction == 6 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			newPosition := processOpcodeFiveSix(instruction, numberOne, numberTwo, position)
			increase = newPosition - position
		}
		if instruction == 7 || instruction == 8 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			out := intcodes[position+3]
			output := processOpcodeSevenEight(instruction, numberOne, numberTwo)
			intcodes[out] = output
			increase = 4

		}
	}
	fmt.Println("Error?")
	return 0, nil
}

func processInstruction(instruction int) (int, bool, bool) {
	instr := strconv.Itoa(instruction)
	e, err := strconv.Atoi(string(instr[len(instr) -1]))
	check(err)
	if len(instr) < 3 {
		if e == 9 {
			e2, err := strconv.Atoi(string(instr[len(instr) -2]))
			check(err)
			if e2 == 9 {
				return 99, false, false
			}
		} else {
			return e, false, false
		}
	}
	modeOne, err := strconv.Atoi(string(instr[len(instr) -3]))
	check(err)
	modeTwo := 0
	if len(instr) == 4 {
		modeTwo, err = strconv.Atoi(string(instr[len(instr) -4]))
		check(err)
		return e, modeOne == 1, modeTwo == 1
	}
	if len(instr) > 4 {
		return instruction, false, false
	}
	check(err)
	return e, modeOne == 1, modeTwo == 1
}

func processOneTwoInstruction(instruction int, numberOne int, numberTwo int) (int, error) {
	if instruction == 1 {
		return numberOne + numberTwo, nil
	}
	if instruction == 2 {
		return numberOne * numberTwo, nil
	}
	return 0, errors.New("Unknown instruction")
}

func processOpcodeThree(intcodes []int, param int, input int) {
	intcodes[param] = input
}

func processOpcodeFour(intcodes []int, param int) int {
	return intcodes[param]
}

func processOpcodeFiveSix(instruction int, numberOne int, numberTwo int, position int) int {
	if instruction == 5 {
		if numberOne != 0 {
			return numberTwo
		}
	}
	if instruction == 6 {
		if numberOne == 0 {
			return numberTwo
		}
	}
	return position + 3
}

func processOpcodeSevenEight(instruction int, numberOne int, numberTwo int) int {
	if instruction == 7 {
		if numberOne < numberTwo {
			return 1
		}
	}
	if instruction == 8 {
		if numberOne == numberTwo {
			return 1
		}
	}
	return 0
}

func processOpcodesFeedBack(codes []int, inChan <-chan int, outChan chan<- int) {
	intcodes := codes
	increase := 4
	instructionMap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
		6: 6,
		7: 7,
		8: 8,
	}
	for position := 0; position < len(intcodes); position += increase {
		instr := intcodes[position]
		instruction, immediateOne, immediateTwo := processInstruction(instr)
		if _, ok := instructionMap[instruction]; !ok || instruction == 99 {
			close(outChan)
			return
			break
		}
		var numberOne int
		var numberTwo int
		if instruction == 1 || instruction == 2 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			out := intcodes[position+3]
			output, err := processOneTwoInstruction(instruction, numberOne, numberTwo)
			check(err)
			intcodes[out] = output
			increase = 4
		}
		if instruction == 3 {
			numberOne = intcodes[position+1]
			input := <- inChan
			processOpcodeThree(intcodes, numberOne, input)
			increase = 2
		}
		if instruction == 4 {
			numberOne = intcodes[position+1]
			res := processOpcodeFour(intcodes, numberOne)
			outChan <- res
			increase = 2
		}
		if instruction == 5 || instruction == 6 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			newPosition := processOpcodeFiveSix(instruction, numberOne, numberTwo, position)
			increase = newPosition - position
		}
		if instruction == 7 || instruction == 8 {
			if immediateOne {
				numberOne = intcodes[position+1]
			} else {
				numberOne = intcodes[intcodes[position+1]]
			}
			if immediateTwo {
				numberTwo = intcodes[position+2]
			} else {
				numberTwo = intcodes[intcodes[position+2]]
			}
			out := intcodes[position+3]
			output := processOpcodeSevenEight(instruction, numberOne, numberTwo)
			intcodes[out] = output
			increase = 4

		}
	}
}