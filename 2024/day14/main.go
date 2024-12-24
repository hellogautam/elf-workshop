package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	gridX   int = 101
	gridY   int = 103
	seconds int = 100
)

func GetInputsFromFile() ([][]int, [][]int) {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	inputLines := strings.Split(string(data), "\n")

	var positions, velocities [][]int

	for _, i := range inputLines {
		posVelStr := strings.Split(i, " ")
		posStr := strings.Split(posVelStr[0], "=")
		posXY := strings.Split(posStr[1], ",")
		posX, err := strconv.Atoi(posXY[0])
		if err != nil {
			panic("code pooped conversion for position x")
		}
		posY, err := strconv.Atoi(posXY[1])
		if err != nil {
			panic("code pooped conversion for position y")
		}
		positions = append(positions, []int{posX, posY})

		velStr := strings.Split(posVelStr[1], "=")
		velXY := strings.Split(velStr[1], ",")
		velX, err := strconv.Atoi(velXY[0])
		if err != nil {
			panic("code pooped conversion for position x")
		}
		velY, err := strconv.Atoi(velXY[1])
		if err != nil {
			panic("code pooped conversion for position y")
		}
		velocities = append(velocities, []int{velX, velY})
	}

	return positions, velocities
}

func simulateMovement(pos []int, vel []int, seconds int) []int {

	posX, posY, velX, velY := pos[0], pos[1], vel[0], vel[1]
	for i := 1; i <= seconds; i++ {
		posX += velX
		if posX < 0 {
			posX += gridX
		} else if posX >= gridX {
			posX -= gridX
		}

		posY += velY
		if posY < 0 {
			posY += gridY
		} else if posY >= gridY {
			posY -= gridY
		}
	}
	return []int{posX, posY}
}

func calcSafetyFactor(endPositions [][]int) int {
	var answer int
	answer = 1
	midX := (gridX - 1) / 2
	midY := (gridY - 1) / 2

	var q1, q2, q3, q4 int

	for _, ep := range endPositions {
		if ep[0] == midX || ep[1] == midY {
			continue
		}
		if ep[0] < midX && ep[1] < midY {
			q1 += 1
		} else if ep[0] > midX && ep[1] < midY {
			q2 += 1
		} else if ep[0] < midX && ep[1] > midY {
			q3 += 1
		} else if ep[0] > midX && ep[1] > midY {
			q4 += 1
		}
	}

	answer = q1 * q2 * q3 * q4
	return answer
}

func SolveP1(positions [][]int, velocities [][]int) (int, [][]int) {
	var answer int

	robots := len(positions)

	var endPositions [][]int

	for i := 0; i < robots; i++ {
		endPositions = append(
			endPositions,
			simulateMovement(positions[i], velocities[i], seconds),
		)
	}

	answer = calcSafetyFactor(endPositions)

	return answer, endPositions
}

func simulateMovementPerSecondForAllRobos(positions [][]int, velocities [][]int) [][]int {
	var newPositions [][]int
	for i := range positions {
		newPositions = append(newPositions, simulateMovement(positions[i], velocities[i], 1))
	}

	return newPositions
}

func markGrid(grid [][]string, positions [][]int) [][]string {
	for _, ep := range positions {
		grid[ep[1]][ep[0]] = "*"
	}

	return grid
}

func printSimulatedGrid(grid [][]string) {
	for _, line := range grid {
		fmt.Println(strings.Join(line, ""))
	}
}

func getNewGrid() (grid [][]string) {
	for y := 0; y < gridY; y++ {
		var xLine []string
		for x := 0; x < gridX; x++ {
			xLine = append(xLine, ".")
		}
		grid = append(grid, xLine)
	}

	return grid
}

func findRectangleWithHeightGreaterThan10(grid [][]string) bool {
	rows, cols := len(grid), len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == "*" {
				for height := 1; i+height < rows; height++ {
					if grid[i+height][j] != "*" {
						break
					}

					if height > 10 {
						return true
					}
				}
			}
		}
	}

	return false
}
func SolveP2(endPositions [][]int, positions [][]int, velocities [][]int) int {
	var answer int

	newPositions := positions
	for i := 0; i < 10000; i++ {
		grid := getNewGrid()
		newPositions = simulateMovementPerSecondForAllRobos(newPositions, velocities)

		grid = markGrid(grid, newPositions)
		if findRectangleWithHeightGreaterThan10(grid) {
			fmt.Println("second ->", i+1)
			printSimulatedGrid(grid)
			fmt.Println("########################################################################################")
		}
	}

	return answer
}

func main() {
	fmt.Println("Hello")
	positions, velocities := GetInputsFromFile()
	p1, endPositions := SolveP1(positions, velocities)
	fmt.Println("p1 -> ", p1)
	fmt.Println("p2 -> ", SolveP2(endPositions, positions, velocities))
}
