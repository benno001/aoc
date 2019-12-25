package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	// "sync"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const positionMode, immediateMode, relativeMode = "0", "1", "2"

type packet struct {
	receiver int
	x        int
	y        int
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
	programs := make([][]int, 50)
	queues := make([][]packet, 50)
	for i := 0; i < 50; i++ {
		intc := make([]int, 10000)
		copy(intc[0:], intcodes)
		programs[i] = intc
	}
	runPrograms(programs, queues)
}

// func monitorProgram(wg *sync.WaitGroup, queues [][]packet) {
// 	wg.Wait()
// 	for _, queue := range queues {
// 		close(queue)
// 	}
// }

func runPrograms(programs [][]int, queues [][]packet) {
	// wg := &sync.WaitGroup{}
	output := make(chan packet)
	// var p packet
	for i := 0; i < 50; i++ {
		// fmt.Println(programs[i])
		processOpcodes(programs[i], queues, i, output)
		// wg.Add(1)
		// o := <-output
		// p = o
		// var o packet

		// select {
		// case o = <-output:
		// 	p = o
		// default:
		// 	fmt.Println("Nothing here")
		// }

	}
	// go monitorProgram(wg, queues)
	// go monitorProgram(wg, output)
	// for i := range output {
	// 	fmt.Println(i)
	// }
	for p := range output {
		fmt.Println(p)
	}
	// p = <-output
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

func processOpcodes(intcodes []int, queues [][]packet, computer int, outChan chan packet) {
	// defer wg.Done()
	// queue := queues[computer]
	relativeBase := 0
	increase := 4
	var out packet
	outPos := 0
	inPos := 0
	bothValuesRead := true
	// var p packet
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
				// fmt.Println(parameterMode[pos], intcodes[position])
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
			time.Sleep(1000)
			if bothValuesRead {
				q := queues[computer]
				if len(q)-1 == inPos || len(q) == 0 {
					fmt.Println("Putting in -1", inPos)
					putParameter(1, -1)
				} else {
					fmt.Println(q)
					// fmt.Println("Putting in", q[inPos].x)
					putParameter(1, q[inPos].x)
					bothValuesRead = false
				}
			} else {
				q := queues[computer]
				fmt.Println(q)
				// fmt.Println("Putting in", q[inPos].y)
				putParameter(1, q[inPos].y)
				bothValuesRead = true
				inPos++
			}
			increase = 2
		}
		if instruction == 4 {
			output := getParameter(1)
			if outPos == 1 {
				// fmt.Println("sending")
				out = packet{out.receiver, output, 0}
				outPos++
			} else if outPos == 2 {
				out = packet{out.receiver, out.x, output}
				fmt.Println("sending", out)
				if out.receiver == 255 {
					outChan <- out
				}
				queues[out.receiver] = append(queues[out.receiver], out)
				outPos = 0
			} else {
				// fmt.Println("sending")
				out = packet{output, 0, 0}
				outPos++
			}
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
}
