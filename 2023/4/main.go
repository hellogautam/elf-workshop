package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	pattern = `\d+`
)

func SolveP1(inputRows []string) int {
	res := 0

	winNums, scratchNums := processInput(inputRows)
	for i := 0; i < len(winNums); i++ {
		points := 0
		for j := 0; j < len(scratchNums[i]); j++ {
			for k := 0; k < len(winNums[i]); k++ {
				if scratchNums[i][j] == winNums[i][k] {
					if points > 0 {
						points = 2 * points
					} else {
						points = 1
					}
				}
			}
		}

		res += points
	}

	return res
}

func processInput(inputRows []string) ([][]int, [][]int) {
	winningNumbers := [][]int{}
	scratchedNumbers := [][]int{}
	regObj := regexp.MustCompile(pattern)

	for _, row := range inputRows {
		wNum, sNum := []int{}, []int{}
		rowSplit := strings.Split(row, ":")
		numberSplit := strings.Split(rowSplit[1], "|")
		winningMatchIndexes := regObj.FindAllStringSubmatchIndex(numberSplit[0], -1)
		for _, ind := range winningMatchIndexes {
			val, _ := strconv.Atoi(numberSplit[0][ind[0]:ind[1]])
			wNum = append(wNum, val)

		}

		scratchedNumbersIndexes := regObj.FindAllStringSubmatchIndex(numberSplit[1], -1)
		for _, ind := range scratchedNumbersIndexes {
			val, _ := strconv.Atoi(numberSplit[1][ind[0]:ind[1]])
			sNum = append(sNum, val)

		}
		// fmt.Println(winningMatchIndexes, scratchedNumbersIndexes)
		// fmt.Println(wNum, sNum)

		winningNumbers = append(winningNumbers, wNum)
		scratchedNumbers = append(scratchedNumbers, sNum)

	}

	return winningNumbers, scratchedNumbers
}

func SolveP2(inputRows []string) int {
	res := 0

	cardCountMap := map[int]int{}
	for i := 0; i < len(inputRows); i++ {
		cardCountMap[i] = 1
	}

	winNums, scratchNums := processInput(inputRows)
	for i := 0; i < len(winNums); i++ {
		count := 0
		for j := 0; j < len(scratchNums[i]); j++ {
			for k := 0; k < len(winNums[i]); k++ {
				if scratchNums[i][j] == winNums[i][k] {
					count += 1
				}
			}
		}

		for c := i + 1; c < i+count+1; c++ {
			cardCountMap[c] = cardCountMap[c] + (cardCountMap[i] * 1)
		}

	}

	for _, v := range cardCountMap {
		res += v
	}

	return res
}

func main() {
	fmt.Println()
	inputRows := GetInputsFromFile()

	fmt.Println("P1 ->", SolveP1(inputRows))
	fmt.Println("P2 ->", SolveP2(inputRows))

}

func GetInputsFromFile() []string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	inputRows := strings.Split(string(data), "\n")

	return inputRows
}
