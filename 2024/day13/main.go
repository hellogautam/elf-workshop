package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GetInputsFromFile() [][]string {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputLines := strings.Split(string(data), "\n")

	var dataSets [][]string

	for i := 0; i < len(inputLines); i += 4 {
		var set []string
		for j := i; j < i+3; j++ {
			set = append(set, inputLines[j])
		}
		dataSets = append(dataSets, set)
	}

	fmt.Println(dataSets)

	return dataSets
}

func splitDetails(inp string, separator string) (result []int) {
	aSplit := strings.Split(inp, ": ")
	aSplitXY := strings.Split(aSplit[1], ", ")
	fmt.Println("xy split", aSplitXY)
	aX := strings.Split(aSplitXY[0], separator)
	aXInt, err := strconv.Atoi(aX[1])
	if err != nil {
		panic("code pooped while getting details")
	}
	aY := strings.Split(aSplitXY[1], separator)
	aYInt, err := strconv.Atoi(aY[1])
	if err != nil {
		panic("code pooped while getting details")
	}
	result = append(result, aXInt, aYInt)

	return
}

func getSetDetails(inp []string) (a []int, b []int, prize []int) {
	// get a values
	a = splitDetails(inp[0], "+")

	b = splitDetails(inp[1], "+")

	prize = splitDetails(inp[2], "=")

	fmt.Println("a->", a, "b->", b, "prize->", prize)

	return
}

func SolveP1(inp [][]string) int {
	var answer int

	for _, i := range inp {
		a, b, prize := getSetDetails(i)
		ax := a[0]
		ay := a[1]
		bx := b[0]
		by := b[1]
		px := 10000000000000 + prize[0]
		py := 10000000000000 + prize[1]

		var na, nb int

		if ((px*ay)-(py*ax))%((ay*bx)-(by*ax)) == 0 {
			nb = ((px * ay) - (py * ax)) / ((ay * bx) - (by * ax))
			if ((px - (bx * nb)) % ax) == 0 {
				na = (px - (bx * nb)) / ax
				fmt.Println("na->", na, "nb", nb)
				answer += na*3 + nb
			}
		}
	}

	return answer
}

func SolveP2() int {
	var answer int

	return answer

}

func main() {
	fmt.Println("Hello")
	data := GetInputsFromFile()
	fmt.Println("p1 -> ", SolveP1(data))
	// fmt.Println("p2 -> ", SolveP2(data))
}
