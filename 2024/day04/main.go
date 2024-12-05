package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

const xmasExpString = `XMAS`
const revXmasExpString = `SAMX`
const enableExpString = `do\(\)`

func getHorizontal(inp string) int {
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

func getEnableCommands(inp string) [][]int {
	regexObj, err := regexp.Compile(enableExpString)
	if err != nil {
		panic(err)
	}

	return regexObj.FindAllStringIndex(inp, -1)
}

func SolveP1(inp []string) int {
	var answer int
	return answer
}

func SolveP2(inp string) int {
	var answer int

	return answer
}

func GetInputsFromFile() []string {
	data, err := os.ReadFile("input_sample.txt")
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
}
