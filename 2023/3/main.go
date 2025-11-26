package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	pattern = `\d+`
)

func SolveP1(inputRows []string) int {
	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}, {1, 1}, {-1, -1}}
	res := 0

	regObj := regexp.MustCompile(pattern)
	matrixMatchIndexes := [][][]int{}
	for _, r := range inputRows {
		matchIndexes := regObj.FindAllStringSubmatchIndex(r, -1)
		matrixMatchIndexes = append(matrixMatchIndexes, matchIndexes)
	}

	for i, indexes := range matrixMatchIndexes {
		for _, ind := range indexes {
			for k := ind[0]; k < ind[1]; k++ {
				x, y := i, k
				for _, d := range dirs {
					nx, ny := x+d[0], y+d[1]
					if nx >= 0 && nx < len(inputRows) && ny >= 0 && ny < len(inputRows[i]) {
						if inputRows[nx][ny] != '.' && !unicode.IsDigit(rune(inputRows[nx][ny])) {
							val, _ := strconv.Atoi(inputRows[i][ind[0]:ind[1]])
							res += val
							k = ind[1]
							break
						}
					}
				}
			}
		}
	}

	return res
}

func SolveP2(inputRows []string) int {
	res := 0

	dirs := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}, {-1, 1}, {1, -1}, {1, 1}, {-1, -1}}

	regObj := regexp.MustCompile(pattern)
	matrixMatchIndexes := [][][]int{}
	for _, r := range inputRows {
		matchIndexes := regObj.FindAllStringSubmatchIndex(r, -1)
		matrixMatchIndexes = append(matrixMatchIndexes, matchIndexes)
	}

	starMap := map[[2]int][]int{}
	for i, indexes := range matrixMatchIndexes {
		for _, ind := range indexes {
			for k := ind[0]; k < ind[1]; k++ {
				x, y := i, k
				for _, d := range dirs {
					nx, ny := x+d[0], y+d[1]
					if nx >= 0 && nx < len(inputRows) && ny >= 0 && ny < len(inputRows[i]) {
						if inputRows[nx][ny] == '*' {
							val, _ := strconv.Atoi(inputRows[i][ind[0]:ind[1]])
							starMap[[2]int{nx, ny}] = append(starMap[[2]int{nx, ny}], val)
							k = ind[1]
							break
						}
					}
				}
			}
		}
	}

	for _, v := range starMap {
		if len(v) == 2 {
			res = res + (v[0] * v[1])
		}
	}

	return res
}

func main() {
	fmt.Println("Hello")

	inputRows := GetInputsFromFile()
	inputMatrix := getInputMatrix(inputRows)
	fmt.Println(inputMatrix)

	fmt.Println("P1 ->", SolveP1(inputRows))
	fmt.Println("P2 ->", SolveP2(inputRows))
}

func getInputMatrix(inputRows []string) [][]string {
	res := [][]string{}
	for _, row := range inputRows {
		res = append(res, strings.Split(row, ""))
	}

	return res
}

func GetInputsFromFile() []string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	inputRows := strings.Split(string(data), "\n")
	fmt.Println(len(inputRows))

	return inputRows
}
