package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SolveP1(inputRows string) int {
	res := 0
	ranges := getRangesFromRows(inputRows)

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			numStr := strconv.Itoa(i)
			if len(numStr)%2 == 0 && numStr[:len(numStr)/2] == numStr[len(numStr)/2:] {
				res += i
			}
		}
	}
	return res
}

func getRangesFromRows(inputRows string) [][]int {
	res := [][]int{}

	splits := strings.Split(inputRows, ",")
	for _, s := range splits {
		rangeSplit := strings.Split(s, "-")
		firstID, _ := strconv.Atoi(rangeSplit[0])
		lastID, _ := strconv.Atoi(rangeSplit[1])

		res = append(res, []int{firstID, lastID})
	}

	return res
}

func SolveP2(inputRows string) int {
	res := 0
	ranges := getRangesFromRows(inputRows)

	for _, r := range ranges {
		for i := r[0]; i <= r[1]; i++ {
			numStr := strconv.Itoa(i)
			if len(numStr)%2 == 0 {
				if numStr[:len(numStr)/2] == numStr[len(numStr)/2:] {
					res += i
					continue
				}
			}
			if len(numStr)%3 == 0 {
				first := numStr[:len(numStr)/3]
				repeats := true
				count := 0
				for i := len(first); i < len(numStr); i = i + len(first) {
					if numStr[i:i+len(first)] == first {
						count++
						continue
					} else {
						repeats = false
						break
					}
				}
				if repeats && count == 2 {
					fmt.Println(i)

					res += i
					continue
				}
			}

			if len(numStr)%5 == 0 {
				first := numStr[:len(numStr)/5]
				repeats := true
				count := 0
				for i := len(first); i < len(numStr); i = i + len(first) {
					if numStr[i:i+len(first)] == first {
						count++
						continue
					} else {
						repeats = false
						break
					}
				}
				if repeats && count == 4 {
					fmt.Println(i)

					res += i
					continue
				}
			}
			if len(numStr)%7 == 0 {
				first := numStr[:len(numStr)/7]
				repeats := true
				count := 0
				for i := len(first); i < len(numStr); i = i + len(first) {
					if numStr[i:i+len(first)] == first {
						count++
						continue
					} else {
						repeats = false
						break
					}
				}
				if repeats && count == 6 {
					fmt.Println(i)
					res += i
					continue
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

func GetInputsFromFile() string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(data)
}
