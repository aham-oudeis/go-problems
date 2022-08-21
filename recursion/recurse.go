package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	}

	return num * factorial(num-1)
}

func sum(low, high int) int {
	if high < low {
		return high
	}

	return high + sum(low, high-1)
}

func main() {
	fmt.Println(factorial(5))
	fmt.Println("sum of integers from 1 to 5")
	fmt.Println(sum(1, 5))
	fmt.Println("vim-go")
}
