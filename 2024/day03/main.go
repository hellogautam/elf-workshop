package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const regExpString = `mul\([0-9]{1,3}?,[0-9]{1,3}?\)`
const disbaleExpString = `don't\(\)`
const enableExpString = `do\(\)`

func getInstructions(inp string) []string {
	regexpObj, err := regexp.Compile(regExpString)
	if err != nil {
		panic(err)
	}

	return regexpObj.FindAllString(inp, -1)
}

func getInstructionIndexes(inp string) [][]int {
	regexpObj, err := regexp.Compile(regExpString)
	if err != nil {
		panic(err)
	}

	return regexpObj.FindAllStringIndex(inp, -1)
}

func getDisableCommands(inp string) [][]int {
	regexObj, err := regexp.Compile(disbaleExpString)
	if err != nil {
		panic(err)
	}

	return regexObj.FindAllStringIndex(inp, -1)
}

func getEnableCommands(inp string) [][]int {
	regexObj, err := regexp.Compile(enableExpString)
	if err != nil {
		panic(err)
	}

	return regexObj.FindAllStringIndex(inp, -1)
}

func SolveP1(inp string) int {
	var answer int
	instructions := getInstructions(inp)
	fmt.Println(instructions)
	for _, ins := range instructions {
		numberString := ins[4 : len(ins)-1]
		numbersStringSplit := strings.Split(numberString, ",")
		if len(numbersStringSplit[0]) < 1 || len(numbersStringSplit[0]) > 3 ||
			len(numbersStringSplit[1]) < 1 || len(numbersStringSplit[1]) > 3 {
			continue
		}
		number0, err := strconv.Atoi(numbersStringSplit[0])
		if err != nil {
			panic(err)
		}
		number1, err := strconv.Atoi(numbersStringSplit[1])
		if err != nil {
			panic(err)
		}
		answer += number0 * number1

	}
	return answer
}

func checkIFInArray(arr []int, el int) (ind int, val int, present bool) {
	for i, b := range arr {
		if b == el {
			return i, b, true
		}
	}

	return 0, 0, false
}

func SolveP2(inp string) int {
	var answer int
	instructions := getInstructionIndexes(inp)
	fmt.Println(instructions)
	disableCmds := getDisableCommands(inp)
	fmt.Println("index disable cmds", disableCmds)
	enableCmds := getEnableCommands(inp)
	fmt.Println("index enable cmds", enableCmds)

	var enabled bool
	var disableIndexes, enableIndexes, instructionIndexes []int
	for _, ins := range instructions {
		instructionIndexes = append(instructionIndexes, ins[0])
	}
	for _, dis := range disableCmds {
		disableIndexes = append(disableIndexes, dis[0])
	}
	for _, ena := range enableCmds {
		enableIndexes = append(enableIndexes, ena[0])
	}

	enabled = true
	for i, _ := range inp {
		if enabled {
			ind, _, present := checkIFInArray(instructionIndexes, i)
			if present {
				numberString := inp[instructions[ind][0]:instructions[ind][1]]

				fmt.Println("debug poop num string", numberString)

				numberString = numberString[4 : len(numberString)-1]
				numbersStringSplit := strings.Split(numberString, ",")
				if len(numbersStringSplit[0]) < 1 || len(numbersStringSplit[0]) > 3 ||
					len(numbersStringSplit[1]) < 1 || len(numbersStringSplit[1]) > 3 {
					continue
				}
				number0, err := strconv.Atoi(numbersStringSplit[0])
				if err != nil {
					panic(err)
				}
				number1, err := strconv.Atoi(numbersStringSplit[1])
				if err != nil {
					panic(err)
				}
				answer += number0 * number1

				fmt.Println("debug poop", answer)

				continue
			} else {
				_, _, present := checkIFInArray(disableIndexes, i)
				if present {
					enabled = false
					continue
				}
			}
		} else {
			_, _, present := checkIFInArray(enableIndexes, i)
			if present {
				enabled = true
			}
		}
	}

	return answer
}

func GetInputsFromFile(fname string) string {
	data, err := os.ReadFile(fname)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func main() {
	fmt.Println("Hello")

	// fmt.Println("sample", SolveP1(GetInputsFromFile("input_sample.txt")))
	// fmt.Println("p1", SolveP1(GetInputsFromFile("input.txt")))
	fmt.Println("samplep2", SolveP2(GetInputsFromFile("input.txt")))
}
