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

		power := gameResults["red"] * gameResults["green"] * gameResults["blue"]
		sum += power

	}
	fmt.Println(sum)
}
