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

	digitMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for sc.Scan() {
		digits := []int{}
		str := sc.Text()
		for i, char := range str {
			if digit, err := strconv.Atoi(string(char)); err == nil {
				digits = append(digits, digit)
			} else {
				for word, number := range digitMap {
					if strings.HasPrefix(str[i:], word) {
						digits = append(digits, number)
						break
					}
				}
			}
		}
		fmt.Println(digits)
		if len(digits) > 0 {
			firstDigit, lastDigit := digits[0], digits[len(digits)-1]
			sum += firstDigit*10 + lastDigit
		} else {
			fmt.Println("No digits found")
		}
	}

	fmt.Println(sum)
}
