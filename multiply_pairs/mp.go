package main

import (
	"fmt"
	"sort"
)

func transform(list []int, change func(val int) int) []int {
	collection := make([]int, len(list))

	for i := range list {
		collection[i] = change(list[i])
	}

	return collection
}

func multiplyAllPairs(l1, l2 []int) []int {
	col := make([]int, 0, len(l1)*len(l2))

	for key := range l1 {
		list := transform(l2, func(num int) int {
			return num * l1[key]
		})

		col = append(col, list...)
	}

	sort.Ints(col)
	return col
}

func main() {
	l1 := []int{12, 3, 4, 5}
	l2 := []int{6, 2, 9, 7}
	fmt.Println(multiplyAllPairs(l1, l2))
	fmt.Println("vim-go")
}
