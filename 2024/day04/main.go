package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const xmasExpString = `XMAS`
const revXmasExpString = `SAMX`

const masExpString = `MAS`
const samExpString = `SAM`

func getMatchesFromLine(inp string) int {
	regexpObj, err := regexp.Compile(xmasExpString)
	if err != nil {
		panic(err)
	}
	revRegexpObj, err := regexp.Compile(revXmasExpString)
	if err != nil {
		panic(err)
	}

	return len(regexpObj.FindAllString(inp, -1)) + len(revRegexpObj.FindAllString(inp, -1))
}

func getInverseInput(inp []string) []string {
	var result []string
	cols := len(inp[0])

	for i := 0; i < cols; i++ {
		var res string
		for _, in := range inp {
			res += string(in[i])
		}
		result = append(result, res)
	}

	fmt.Println("inverse", result)

	return result
}

func getDiagonals(inp []string) []string {
	var diagonals []string
	extractDiagonal := func(matrix []string, startRow, startCol, rowInc, colInc int) string {
		var diagonal string
		for row, col := startRow, startCol; row >= 0 && col >= 0 && row < len(matrix) && col < len(matrix[0]); row, col = row+rowInc, col+colInc {
			diagonal += string(matrix[row][col])
		}
		return diagonal
	}

	// Extract left-to-right diagonals
	for i := 0; i < len(inp); i++ {
		// Diagonals starting from the top row
		diagonals = append(diagonals, extractDiagonal(inp, i, 0, 1, 1))
	}

	for j := 1; j < len(inp[0]); j++ {
		// Diagonals starting from the first column
		diagonals = append(diagonals, extractDiagonal(inp, 0, j, 1, 1))
	}

	// Extract right-to-left diagonals
	for i := 0; i < len(inp); i++ {
		// Diagonals starting from the top row
		diagonals = append(diagonals, extractDiagonal(inp, i, len(inp[0])-1, 1, -1))
	}

	for j := len(inp[0]) - 2; j >= 0; j-- {
		// Diagonals starting from the last column
		diagonals = append(diagonals, extractDiagonal(inp, 0, j, 1, -1))
	}

	return diagonals
}

func SolveP1(inp []string) int {
	var answer int

	inverseInputs := getInverseInput(inp)

	for _, str := range inp {
		answer += getMatchesFromLine(str)
	}
	for _, inStr := range inverseInputs {
		answer += getMatchesFromLine(inStr)
	}

	diagonals := getDiagonals(inp)
	for _, diag := range diagonals {
		answer += getMatchesFromLine(diag)
	}

	fmt.Println(answer)

	return answer
}

func SolveP2(inp []string) int {
	var answer int

	for i := 1; i < len(inp)-1; i++ {
		for j := 1; j < len(inp[i])-1; j++ {
			if string(inp[i][j]) == "A" {
				ltr := string(inp[i-1][j-1]) + string(inp[i][j]) + string(inp[i+1][j+1])
				rtl := string(inp[i-1][j+1]) + string(inp[i][j]) + string(inp[i+1][j-1])
				if (ltr == "MAS" || ltr == "SAM") && (rtl == "MAS" || rtl == "SAM") {
					answer += 1
				}
			}
		}
	}

	return answer
}

func GetInputsFromFile() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "\n")
}

func main() {
	fmt.Println("Hello")
	rows := GetInputsFromFile()
	fmt.Println(rows)
	fmt.Println("p1 -> ", SolveP1(rows))
	fmt.Println("p2 -> ", SolveP2(rows))
}
