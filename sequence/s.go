package main

import "fmt"

func sequence(total, start int) []int {
	col := []int{}

	for i := 0; i < total; i++ {
		col = append(col, start)
		start++
	}

	return col
}
func main() {
	fmt.Println("vim-go")
	fmt.Println(sequence(4, -7))
}
