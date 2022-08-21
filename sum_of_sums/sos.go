package main

import "fmt"

func sumOfSums(list []int) int {
	var total int
	var current int

	for idx := range list {
		current += list[idx]
		total += current
	}

	return total
}

func main() {
	fmt.Println("vim-go")
	l1 := []int{3, 5, 2}
	l2 := []int{1, 5, 7, 3}
	fmt.Println(sumOfSums(l1))
	fmt.Println(sumOfSums(l2))
}
