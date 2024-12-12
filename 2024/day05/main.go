package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInputsFromFile() (map[int][]int, [][]int) {
	dataSection1, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var rulesString []string
	// var rules map[int][]int
	rules := make(map[int][]int)
	rulesString = strings.Split(string(dataSection1), "\n")
	fmt.Println(rulesString)
	for _, rule := range rulesString {
		ruleSplit := strings.Split(rule, "|")
		if len(ruleSplit) == 2 {
			key, err := strconv.Atoi(ruleSplit[0])
			if err != nil {
				panic("code poop")
			}
			val, err := strconv.Atoi(ruleSplit[1])
			if err != nil {
				panic("code poop")
			}
			_, ok := rules[key]
			// If the key exists
			if ok {
				rules[key] = append(rules[key], val)
			} else {
				rules[key] = []int{}
				rules[key] = append(rules[key], val)
			}

			// fmt.Println("key", key, "value", val)
		} else {
			panic("code poop")
		}
	}

	fmt.Println("rules", rules)

	dataSection2, err := os.ReadFile("input_2.txt")
	if err != nil {
		panic(err)
	}

	var updates [][]int
	updateRows := strings.Split(string(dataSection2), "\n")
	for _, updateStr := range updateRows {
		update := strings.Split(updateStr, ",")
		var updateInts []int
		for _, u := range update {
			val, err := strconv.Atoi(u)
			if err != nil {
				panic("code poop")
			}
			updateInts = append(updateInts, val)
		}
		updates = append(updates, updateInts)
	}

	fmt.Println("updates", updates)

	return rules, updates
}

func checkIFInArray(arr []int, el int) (ind int, val int, present bool) {
	for i, b := range arr {
		if b == el {
			return i, b, true
		}
	}

	return 0, 0, false
}

func SolveP1(rules map[int][]int, updates [][]int) int {
	var answer int
	var incorrectUpdates [][]int

	for _, update := range updates {
		incorrect := false
		for j, u := range update {
			_, ok := rules[u]
			// If the key exists
			if ok {
				for k := j - 1; k >= 0; k-- {
					// fmt.Println("i", i, "j", j)
					_, _, present := checkIFInArray(rules[u], update[k])
					if present {
						incorrect = true
						break
					}
				}
			} else {
				continue
			}
		}

		if !incorrect {
			mid := update[len(update)/2]
			fmt.Println("mid", mid)
			answer += mid
		} else {
			incorrectUpdates = append(incorrectUpdates, update)
		}

	}

	SolveP2(rules, incorrectUpdates)

	return answer
}

func insertInt(array []int, value int, index int) []int {
	return append(array[:index], append([]int{value}, array[index:]...)...)
}

func removeInt(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func moveInt(array []int, srcIndex int, dstIndex int) []int {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}

func SolveP2(rules map[int][]int, updates [][]int) {
	var answer int

	for _, u := range updates {
		resArr := u
		for uIndex, uElement := range u {
			leastIndex := -1
			_, ok := rules[uElement]
			// If the key exists
			if ok {
				for k := uIndex - 1; k >= 0; k-- {
					_, _, present := checkIFInArray(rules[uElement], u[k])
					if present {
						leastIndex = k
					}
				}
			}
			if leastIndex != -1 {
				resArr = moveInt(resArr, uIndex, leastIndex)
			}
		}
		mid := resArr[len(resArr)/2]
		fmt.Println("mid", mid)
		answer += mid
	}

	fmt.Println("p2 -> ", answer)
}

func main() {
	fmt.Println("Hello")
	rules, updates := GetInputsFromFile()
	// fmt.Println(rows)
	fmt.Println("p1 -> ", SolveP1(rules, updates))
	// fmt.Println("p2 -> ", SolveP2(rows))
}
