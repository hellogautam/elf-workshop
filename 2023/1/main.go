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
	numPattern = "one|two|three|four|five|six|seven|eight|nine|[1-9]"
)

func SolveP1(inputRows []string) int {
	sum := 0

	for _, row := range inputRows {
		nums := []rune{}
		for _, c := range row {
			if unicode.IsDigit(c) {
				nums = append(nums, c)
			}
		}
		var numStr string
		if len(nums) > 0 {
			numStr = string(nums[0]) + string(nums[len(nums)-1])
		} else {
			continue
		}

		numVal, _ := strconv.Atoi(numStr)
		sum += numVal
	}

	return sum
}

// Function to find all potentially overlapping matches in a string
func findAllOverlappingMatches(input string, pattern string) []string {
	// Compile the regex pattern
	re := regexp.MustCompile(pattern)

	var matches []string
	searchIndex := 0

	for searchIndex < len(input) {
		matchIndices := re.FindStringSubmatchIndex(input[searchIndex:])

		if matchIndices == nil {
			break
		}

		absoluteStartIndex := searchIndex + matchIndices[0]
		absoluteEndIndex := searchIndex + matchIndices[1]

		match := input[absoluteStartIndex:absoluteEndIndex]
		matches = append(matches, match)

		searchIndex = absoluteStartIndex + 1
	}

	return matches
}

func SolveP2(inputRows []string) int {

	sum := 0

	for _, row := range inputRows {
		// nums := regexObj.FindAllString(row, -1)
		nums := findAllOverlappingMatches(row, numPattern)
		fmt.Println("num matches", nums)
		numStr := getNumMapping(nums[0]) + getNumMapping(nums[len(nums)-1])
		numVal, _ := strconv.Atoi(numStr)
		sum += numVal
	}

	return sum
}

func main() {
	fmt.Println("Hello World!")

	inputRows := GetInputsFromFile()

	fmt.Println(inputRows)

	fmt.Println("P1 ->", SolveP1(inputRows))
	fmt.Println("P2 ->", SolveP2(inputRows))

}

func getNumMapping(s string) string {
	switch s {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return s
	}
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
