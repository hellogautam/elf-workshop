package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolveP1(ral [][]int) int {
	var answer int

	for _, report := range ral {
		if checkIfSafe(report) {
			answer += 1
		}
	}
	return answer
}

func checkIfSafe(report []int) bool {
	safe := true
	order := "neutral"
	for i := 0; i < len(report)-1; i++ {
		if order == "neutral" {
			if report[i] < report[i+1] {
				order = "increasing"
			} else if report[i] > report[i+1] {
				order = "decreasing"
			} else {
				safe = false
			}
		}

		switch order {
		case "increasing":
			if report[i+1]-report[i] <= 3 && report[i+1]-report[i] >= 1 {
				continue
			} else {
				safe = false
			}

		case "decreasing":
			if report[i]-report[i+1] <= 3 && report[i]-report[i+1] >= 1 {
				continue
			} else {
				safe = false
			}

		}

		if !safe {
			break
		}
	}

	return safe
}

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func checkIfSafeByRemovingLevels(report []int) bool {
	var safe bool

	fmt.Println("report", report)
	for i := 0; i < len(report); i++ {
		copyReport := make([]int, len(report))
		copy(copyReport, report)
		// fmt.Println("report levels removed", deleteElement(copyReport, i))
		safe = checkIfSafe(deleteElement(copyReport, i))
		if safe {
			break
		}
	}

	return safe
}

func SolveP2(ral [][]int) int {
	var answer int

	for _, report := range ral {
		if checkIfSafe(report) {
			answer += 1
			continue
		}
		if checkIfSafeByRemovingLevels(report) {
			answer += 1
			continue
		}

	}
	return answer
}

func GetInputsFromFile() [][]int {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var reportAndLevels [][]int

	inputRows := strings.Split(string(data), "\n")

	for _, row := range inputRows {
		levels := strings.Split(row, " ")
		var levelIntList []int
		for _, level := range levels {
			levelInt, err := strconv.Atoi(string(level))
			if err != nil {
				panic(err)
			}
			levelIntList = append(levelIntList, levelInt)
		}
		reportAndLevels = append(reportAndLevels, levelIntList)
	}

	// fmt.Println(reportAndLevels)

	return reportAndLevels
}

func main() {
	fmt.Println("Hello")
	reportAndLevels := GetInputsFromFile()
	fmt.Println(SolveP1(reportAndLevels))
	fmt.Println(SolveP2(reportAndLevels))

}
