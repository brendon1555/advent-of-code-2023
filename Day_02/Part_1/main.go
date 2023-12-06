package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	input, err := os.Open("../input.txt")
	check(err)
	defer input.Close()

	sc := bufio.NewScanner(input)

	sum := 0
	for sc.Scan() {
		gameResults := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}

		str := sc.Text()
		game := strings.Split(str, ": ")
		index, err := strconv.Atoi(strings.Split(game[0], " ")[1])
		check(err)
		allSets := game[1]
		sets := strings.Split(allSets, "; ")
		for _, set := range sets {
			colors := strings.Split(set, ", ")
			for _, color := range colors {
				colorScore := strings.Split(color, " ")
				score, err := strconv.Atoi(colorScore[0])
				check(err)
				colorName := colorScore[1]
				gameResults[colorName] = max(gameResults[colorName], score)
			}
		}

		if gameResults["red"] <= 12 && gameResults["green"] <= 13 && gameResults["blue"] <= 14 {
			sum += index
		}
	}
	fmt.Println(sum)
}
