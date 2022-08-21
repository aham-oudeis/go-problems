package main

import (
	"fmt"
	"strconv"
)

func transform(list []rune, change func(val rune) int) []int {
	collection := make([]int, len(list))

	for i := range list {
		collection[i] = change(list[i])
	}

	return collection
}

func total(list []int) int {
	var sum int

	for i := range list {
		sum += list[i]
	}

	return sum
}

func sum(num int) int {
	digits := []rune(strconv.Itoa(num))

	runeToInt := func(r rune) int {
		num, _ := strconv.Atoi(string(r))
		return num
	}

	digitsSlice := transform(digits, runeToInt)

	return total(digitsSlice)
}

func main() {
	fmt.Println(sum(25))
	fmt.Println(sum(253))
	fmt.Println("vim-go")
}
