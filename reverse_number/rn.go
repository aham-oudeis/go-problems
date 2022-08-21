package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(list []string) []string {
	l := len(list)
	last := l - 1
	collection := make([]string, l)

	for i := 0; i < l; i++ {
		collection[last-i] = list[i]
	}

	return collection
}

func reverseNumber(int int) int {
	digits := strings.Split(strconv.Itoa(int), "")

	digitsReverse := reverse(digits)
	reverseStr := strings.Join(digitsReverse, "")
	reversedNum, _ := strconv.Atoi(reverseStr)

	return reversedNum
}

func main() {
	fmt.Println(reverseNumber(123458))
	fmt.Println("vim-go")
}
