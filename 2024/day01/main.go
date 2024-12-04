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
	soretedLoc1 := sortInt(loc1)
	soretedLoc2 := sortInt(loc2)

	fmt.Println(soretedLoc1)
	fmt.Println(soretedLoc2)

	for i, _ := range soretedLoc1 {
		if soretedLoc1[i] > soretedLoc2[i] {
			answer += soretedLoc1[i] - soretedLoc2[i]
		} else {
			answer += soretedLoc2[i] - soretedLoc1[i]
		}
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
	fmt.Println(SolveP1(locationIDs1, locationIDs2))
}
