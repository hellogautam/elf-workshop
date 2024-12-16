package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type direction string

const (
	north direction = "n"
	south direction = "s"
	east  direction = "e"
	west  direction = "w"
)

type pos struct {
	x   int
	y   int
	dir direction
}

func turnRight(inp direction) direction {
	switch inp {
	case north:
		return west
	case south:
		return east
	case west:
		return south
	case east:
		return north
	default:
		panic("code pooped")
	}
}

func GetInputsFromFile() [][]string {
	dataSection1, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	var dataMap [][]string

	dataMapStrings := strings.Split(string(dataSection1), "\n")
	for _, d := range dataMapStrings {
		dataRow := strings.Split(string(d), "")
		dataMap = append(dataMap, dataRow)
	}

	return dataMap
}

func moveGuard(dataMap [][]string, gp pos) ([][]string, pos) {
	switch gp.dir {
	case north:
		if string(dataMap[gp.x-1][gp.y]) == "#" {
			return dataMap, pos{
				x:   gp.x,
				y:   gp.y,
				dir: turnRight(gp.dir),
			}
		}

		currPos := pos{
			x:   gp.x - 1,
			y:   gp.y,
			dir: gp.dir,
		}

		dataMap[currPos.x][currPos.y] = "X"

		return dataMap, currPos
	case south:
		if string(dataMap[gp.x+1][gp.y]) == "#" {
			return dataMap, pos{
				x:   gp.x,
				y:   gp.y,
				dir: turnRight(gp.dir),
			}
		}
		currPos := pos{
			x:   gp.x + 1,
			y:   gp.y,
			dir: gp.dir,
		}
		dataMap[currPos.x][currPos.y] = "X"
		return dataMap, currPos
	case east:
		if string(dataMap[gp.x][gp.y-1]) == "#" {
			return dataMap, pos{
				x:   gp.x,
				y:   gp.y,
				dir: turnRight(gp.dir),
			}
		}
		currPos := pos{
			x:   gp.x,
			y:   gp.y - 1,
			dir: gp.dir,
		}
		dataMap[currPos.x][currPos.y] = "X"
		return dataMap, currPos
	case west:
		if string(dataMap[gp.x][gp.y+1]) == "#" {
			return dataMap, pos{
				x:   gp.x,
				y:   gp.y,
				dir: turnRight(gp.dir),
			}
		}
		currPos := pos{
			x:   gp.x,
			y:   gp.y + 1,
			dir: gp.dir,
		}
		dataMap[currPos.x][currPos.y] = "X"
		return dataMap, currPos
	default:
		panic("code pooped for moving guard")
	}
}

func SolveP1(dataMap [][]string) (int, [][]int, pos) {
	var answer int
	var start pos

	//get starting position
	for i := 0; i < len(dataMap); i++ {
		for j := 0; j < len(dataMap[i]); j++ {
			if string(dataMap[i][j]) == "^" {
				start.x = i
				start.y = j
				start.dir = north
				break
			}
		}
	}

	var guardPos pos = start

	dataMap[guardPos.x][guardPos.y] = "X"

	for guardPos.x > 0 && guardPos.y > 0 && guardPos.x < len(dataMap)-1 && guardPos.y < len(dataMap[0])-1 {
		dataMap, guardPos = moveGuard(dataMap, guardPos)
	}

	var allVisits [][]int

	for x, row := range dataMap {
		for y, d := range row {
			var coordinates []int
			if d == "X" {
				coordinates = append(coordinates, x)
				coordinates = append(coordinates, y)
				answer += 1
				allVisits = append(allVisits, coordinates)
			}
		}
	}

	return answer, allVisits, start
}

func posToString(inp pos) string {
	return strconv.Itoa(inp.x) + "," + strconv.Itoa(inp.y) + "," + string(inp.dir)
}

func SolveP2(allVisits [][]int, start pos) int {
	var answer int

	for _, visit := range allVisits {
		dataMap := GetInputsFromFile()

		if start.x == visit[0] && start.y == visit[1] {
			continue
		}
		dataMap[visit[0]][visit[1]] = "#"

		var guardPos pos = start

		dataMap[guardPos.x][guardPos.y] = "X"

		visitMap := make(map[string]bool)
		visitMap[posToString(guardPos)] = true

		for guardPos.x > 0 && guardPos.y > 0 && guardPos.x < len(dataMap)-1 && guardPos.y < len(dataMap[0])-1 {
			dataMap, guardPos = moveGuard(dataMap, guardPos)
			posStr := posToString(guardPos)
			_, ok := visitMap[posStr]
			// If the key exists
			if ok {
				answer += 1
				break
			} else {
				visitMap[posStr] = true
			}
		}
	}

	return answer
}

func main() {
	fmt.Println("Hello")
	dataMap := GetInputsFromFile()
	answerP1, allVisits, startPos := SolveP1(dataMap)
	fmt.Println("p1 -> ", answerP1)
	fmt.Println("p2 -> ", SolveP2(allVisits, startPos))
}
