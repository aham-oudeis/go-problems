package main

import (
	"fmt"
)

func transform(list []rune) []int {
	var col []int = make([]int, len(list))

	for i := range list {
		col[i] = int(list[i])
	}

	return col
}

func sum(col []int) int {
	var total int = 0

	for _, e := range col {
		total += e
	}

	return total
}

func utf16(str string) int {
	codePoints := []rune(str)
	return sum(transform(codePoints))
}

func main() {
	fmt.Println(utf16("Four score"))
}
