package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isNotSymbol(r rune) bool {
	return (r < '0' || r > '9') && r != '.'
}

func hasSymbols(str string) bool {
	return strings.IndexFunc(str, isNotSymbol) != -1
}

func main() {
	input, err := os.ReadFile("../input.txt")
	check(err)
	lines := strings.Split(string(input), "\n")

	re := regexp.MustCompile(`\d+`)

	sum := 0

	for lineIndex, line := range lines {
		numberIndex := re.FindAllIndex([]byte(line), -1)
		numbers := re.FindAllString(line, -1)

		for i, indexArr := range numberIndex {
			extendedIndex := [2]int{int(math.Max(float64(indexArr[0]-1), 0.0)), int(math.Min(float64(indexArr[1]+1), float64(len(line))))}
			prevLineChars := "."
			nextLineChars := "."
			if lineIndex-1 >= 0 {
				prevLineChars = lines[lineIndex-1][extendedIndex[0]:extendedIndex[1]]
			}
			if lineIndex+1 < len(lines) {
				nextLineChars = lines[lineIndex+1][extendedIndex[0]:extendedIndex[1]]
			}
			if hasSymbols(prevLineChars) || hasSymbols(line[extendedIndex[0]:extendedIndex[1]]) || hasSymbols(nextLineChars) {
				number, err := strconv.Atoi(numbers[i])
				check(err)
				sum += number
			}
		}
	}

	fmt.Println(sum)
}
