package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sortInt(loc []int) []int {
	var temp int
	for i := 0; i < len(loc)-1; i++ {
		small := i
		for j := i + 1; j < len(loc); j++ {
			if loc[small] > loc[j] {
				small = j
			}
		}
		temp = loc[i]
		loc[i] = loc[small]
		loc[small] = temp
	}

	return loc
}

func SolveP1(loc1 []int, loc2 []int) int {
	var answer int

	for i, _ := range loc1 {
		if loc1[i] > loc2[i] {
			answer += loc1[i] - loc2[i]
		} else {
			answer += loc2[i] - loc1[i]
		}
	}

	return answer

}

func SolveP2(loc1 []int, loc2 []int) int {
	var answer int
	answer = 0

	for i := 0; i < len(loc1); i++ {
		similar := 0
		for j := 0; j < len(loc2); j++ {
			if loc1[i] == loc2[j] {
				similar += 1
			}

		}
		answer += loc1[i] * similar
	}

	return answer
}

func GetInputsFromFile() ([]int, []int) {
	var locationIDs1, locationIDs2 []int
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	inputRows := strings.Split(string(data), "\n")
	fmt.Println(len(inputRows))

	for _, row := range inputRows {
		rowSplit := strings.Split(string(row), " ")

		re1, err := strconv.Atoi(string(rowSplit[0]))
		if err != nil {
			panic("poop 1")
		}

		re2, err := strconv.Atoi(rowSplit[len(rowSplit)-1])
		if err != nil {
			panic("poop 2")
		}
		locationIDs1 = append(locationIDs1, re1)
		locationIDs2 = append(locationIDs2, re2)
	}

	return locationIDs1, locationIDs2
}

func main() {
	fmt.Println("Hello")

	locationIDs1, locationIDs2 := GetInputsFromFile()
	sortedLoc1 := sortInt(locationIDs1)
	sortedLoc2 := sortInt(locationIDs2)
	fmt.Println(SolveP1(sortedLoc1, sortedLoc2))
	fmt.Println(SolveP2(sortedLoc1, sortedLoc2))
}
