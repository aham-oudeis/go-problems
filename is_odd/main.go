package main

import (
	"fmt"
	"math"
)

func isOdd(num int) bool {
	absNum := int(math.Abs(float64(num)))
	return absNum%2 == 1
}

func logOdd(num int) {
	for i := 0; i < num; i++ {
		if isOdd(i) {
			fmt.Println(i)
		}
	}
}

func logEven(num int) {
	for i := 0; i < num; i++ {
		if !isOdd(i) {
			fmt.Println(i)
		}
	}
}

type filterFunc = func(n int) bool

func log(num int, test filterFunc) {
	for i := 0; i < num; i++ {
		if test(i) {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println(isOdd(2))
	fmt.Println(isOdd(5))
	fmt.Println(isOdd(-5))
	fmt.Println(isOdd(0))
	fmt.Println(isOdd(-8))

	logOdd(100)
}
