package main

import "fmt"

func sequence(num int) []int {
	col := []int{}

	for i := 1; i <= num; i++ {
		col = append(col, i)
	}

	return col
}

func main() {
	fmt.Println(sequence(5))
	fmt.Println("vim-go")
}
