package main

import "fmt"

func centerOf(str string) string {
	l := len(str)
	half := l / 2

	if l%2 == 0 {
		return str[half-1 : half+1]
	}

	return str[half : half+1]
}

func main() {
	fmt.Println(centerOf("Lanch"))
	fmt.Println(centerOf("Launch"))
	fmt.Println(centerOf("Launch School"))
	fmt.Println("vim-go")
}
