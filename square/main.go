package main

import "fmt"

type Number interface {
	int | float64 | float32
}

func square[T Number](val T) T {
	return val * val
}

func main() {
	fmt.Println(square(5))
	fmt.Println(square(3.2))
	fmt.Println(square(int(73)))
}
