package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolveP1(inputRows []string) int {
	res := 0

	curr := 50

	for _, r := range inputRows {
		val, _ := strconv.Atoi(r[1:])
		if val >= 100 {
			val = val % 100
		}
		if r[0] == 'L' {
			if val > curr {
				curr = 100 - (val - curr)
			} else {
				curr = curr - val
			}
		}
		if r[0] == 'R' {
			curr = curr + val
			if curr >= 100 {
				curr = curr - 100
			}
		}

		if curr == 0 {
			res += 1
		}

	}

	return res
}

func SolveP2(inputRows []string) int {
	res := 0

	curr := 50

	for _, r := range inputRows {
		val, _ := strconv.Atoi(r[1:])

		if r[0] == 'L' {
			for i := 0; i < val; i++ {
				curr--
				if curr == 0 {
					res++
				}

				if curr < 0 {
					curr += 100
				}
			}
		}
		if r[0] == 'R' {
			for i := 0; i < val; i++ {
				curr++
				if curr%100 == 0 {
					res++
				}

				if curr > 100 {
					curr -= 100
				}
			}
		}

	}

	return res
}

func main() {
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
