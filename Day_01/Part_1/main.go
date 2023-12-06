package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	re := regexp.MustCompile(`\d`)
	sum := 0

	for sc.Scan() {
		numArr := re.FindAllString(sc.Text(), -1)
		if len(numArr) > 0 {
			firstNum, lastNum := numArr[0], numArr[len(numArr)-1]
			i, err := strconv.Atoi(firstNum + lastNum)
			check(err)
			sum += i
		}
	}

	fmt.Println(sum)
}
