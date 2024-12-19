package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// file , free space , file, free space, file ...

func GetInputsFromFile() []string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return strings.Split(string(data), "")
}

func expandStorageInfo(inp []string) []string {
	var expStr []string
	var fileID int

	for i, c := range inp {
		cInt, err := strconv.Atoi(string(c))
		if err != nil {
			panic("code pooped exp str 1")
		}
		if i%2 == 0 {
			fileIDStr := strconv.Itoa(fileID)
			for k := 0; k < cInt; k++ {
				expStr = append(expStr, fileIDStr)
			}
			fileID += 1
		} else {
			for k := 0; k < cInt; k++ {
				expStr = append(expStr, ".")
			}
		}
	}

	return expStr
}

func calcChecksum(inp []string) int {
	var answer int

	for i, c := range inp {
		if string(c) != "." {
			cInt, err := strconv.Atoi(string(c))
			if err != nil {
				panic("code pooped checksum int conversionl")
			}
			answer += cInt * i
		}
	}

	return answer
}

func replaceCharacters(str string, index1, index2 int) string {
	result := []byte(str)
	if index1 >= 0 && index1 < len(result) {
		result[index1] = str[index2]
	}
	if index2 >= 0 && index2 < len(result) {
		result[index2] = str[index1]
	}
	return string(result)
}

func shifterP1(inp []string) []string {
	shiftedStr := inp

	for i, j := 0, len(shiftedStr)-1; i < j; {

		if string(shiftedStr[i]) != "." {
			i++
			continue
		}
		if string(shiftedStr[j]) == "." {
			j--
			continue
		}

		// if string(inp[i]) == "." && string(inp[j]) != "." {
		temp := shiftedStr[i]
		shiftedStr[i] = shiftedStr[j]
		shiftedStr[j] = temp

		// shiftedStr = replaceCharacters(shiftedStr, i, j)
		i++
		j--
		// }
	}

	// fmt.Println("shifted str", shiftedStr)

	return shiftedStr
}

func SolveP1(inp []string) int {
	var answer int

	expStr := expandStorageInfo(inp)
	shiftedStr := shifterP1(expStr)

	answer = calcChecksum(shiftedStr)

	return answer
}

// _, ok := visitMap[posStr]
// 			// If the key exists
// 			if ok {
// 				answer += 1
// 				break
// 			} else {
// 				visitMap[posStr] = true
// 			}

func freeSpaceCalculator(inp []string) ([][]int, [][]int) {
	var freeSpaces, storageInfo [][]int
	var free bool
	startIndex := 0

	fmt.Println(inp)

	for i := 0; i < len(inp); i++ {
		if free {
			if "." == inp[i] {
				continue
			} else {
				free = false
				freeSpaces = append(freeSpaces, []int{startIndex, i})
				startIndex = i
			}
		} else {
			if "." == inp[i] {
				free = true
				fileID, err := strconv.Atoi(inp[i-1])
				if err != nil {
					panic("code pooped converting file id")
				}
				storageInfo = append(storageInfo, []int{fileID, startIndex, i})
				startIndex = i
			} else if i != 0 && inp[i] != inp[i-1] {
				fileID, err := strconv.Atoi(inp[i-1])
				if err != nil {
					panic("code pooped converting file id")
				}
				storageInfo = append(storageInfo, []int{fileID, startIndex, i})
				startIndex = i
			} else {
				if i == len(inp)-1 {
					fileID, err := strconv.Atoi(inp[i-1])
					if err != nil {
						panic("code pooped converting file id")
					}
					storageInfo = append(storageInfo, []int{fileID, startIndex, i + 1})
					startIndex = i
				}
				continue
			}
		}

	}

	fmt.Println("free spaces ->", freeSpaces)
	fmt.Println("storage info ->", storageInfo)

	return freeSpaces, storageInfo
}

func replaceChunk(arr []string, freeSpace []int, storageInfo []int, fileLength int) []string {

	for i := freeSpace[0]; i < freeSpace[0]+fileLength; i++ {
		arr[i] = strconv.Itoa(storageInfo[0])
	}

	for j := storageInfo[1]; j < storageInfo[2]; j++ {
		arr[j] = "."
	}

	return arr
}

// freeSpaces map -> key is contiguous space and value is indexes where it occurs
// storageInfo map -> key is fileID and value is number of blocks required
func shifterP2(inp []string, freeSpaces [][]int, storageInfo [][]int) []string {
	shiftedStr := inp

	for i := len(storageInfo) - 1; i >= 0; i-- {
		lengthOfFile := storageInfo[i][2] - storageInfo[i][1]
		for j := 0; j < len(freeSpaces)-1; j++ {
			if storageInfo[i][1] > freeSpaces[j][0] {
				lengthOfFreeSpace := freeSpaces[j][1] - freeSpaces[j][0]
				if lengthOfFreeSpace >= lengthOfFile {
					shiftedStr = replaceChunk(shiftedStr, freeSpaces[j], storageInfo[i], lengthOfFile)
					freeSpaces[j][0] += lengthOfFile
					break
				}
			}
		}
	}

	return shiftedStr
}

func SolveP2(inp []string) int {
	var answer int

	expStr := expandStorageInfo(inp)

	fmt.Println("expanded", expStr)

	freeSpaceMap, storageInfo := freeSpaceCalculator(expStr)
	shiftedStr := shifterP2(expStr, freeSpaceMap, storageInfo)

	fmt.Println(strings.Join(shiftedStr, ""))

	answer = calcChecksum(shiftedStr)

	return answer

}

func main() {
	fmt.Println("Hello")
	data := GetInputsFromFile()
	fmt.Println("p1 -> ", SolveP1(data))
	fmt.Println("p2 -> ", SolveP2(data))
}
