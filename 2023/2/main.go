package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id    int
	blue  int
	green int
	red   int
}

func SolveP1(games []Game, red, green, blue int) int {
	res := 0

	for _, g := range games {
		if g.blue <= blue && g.red <= red && g.green <= green {
			res += g.id
		}
	}

	return res
}

func SolveP2(games []Game) int {
	res := 0

	for _, g := range games {
		res = res + (g.blue * g.green * g.red)
	}

	return res
}

func main() {
	fmt.Println("hello")
	inputRows := GetInputsFromFile()

	games := processInput(inputRows)

	red, green, blue := 12, 13, 14

	fmt.Println("P1 ->", SolveP1(games, red, green, blue))
	fmt.Println("P2 ->", SolveP2(games))
}

func processInput(inputRows []string) []Game {
	res := []Game{}

	for _, row := range inputRows {
		game := Game{}
		maxRed, maxBlue, maxGreen := 0, 0, 0
		gameInfo := strings.Split(row, ": ")
		gameHeader := strings.Split(gameInfo[0], " ")
		game.id, _ = strconv.Atoi(gameHeader[1])

		gamesPlayed := strings.Split(gameInfo[1], "; ")
		for _, g := range gamesPlayed {
			pulled := strings.Split(g, ", ")
			for _, p := range pulled {
				pInfo := strings.Split(p, " ")
				pNum, _ := strconv.Atoi(pInfo[0])
				switch pInfo[1] {
				case "red":
					if maxRed < pNum {
						maxRed = pNum
					}
				case "green":
					if maxGreen < pNum {
						maxGreen = pNum
					}
				case "blue":
					if maxBlue < pNum {
						maxBlue = pNum
					}
				}
			}
		}
		game.blue = maxBlue
		game.red = maxRed
		game.green = maxGreen
		res = append(res, game)

	}

	return res
}

func GetInputsFromFile() []string {
	fileName := "input.txt"
	data, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	inputRows := strings.Split(string(data), "\n")
	fmt.Println(len(inputRows))

	return inputRows
}
