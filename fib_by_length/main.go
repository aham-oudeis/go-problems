package main

import (
	"fmt"
	"strconv"
)

func findFibIndexByLength(num int) int {
	if num == 1 {
		return num
	}

	var count int = 2
	var findFibLength func(int, int, *int) int

	findFibLength = func(next, prev int, index *int) int {
		length := len(strconv.Itoa(next))
		if length == num {
			return *index
		}

		*index = *index + 1
		temp := next
		next = next + prev
		prev = temp

		return findFibLength(next, prev, index)
	}

	return findFibLength(1, 1, &count)
}

func main() {
	fmt.Println(findFibIndexByLength(2))
	fmt.Println(findFibIndexByLength(16))
	fmt.Println(findFibIndexByLength(20))
	//throws an error for number with larger length like more than 20
}
