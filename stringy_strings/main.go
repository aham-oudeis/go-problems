package main

import (
	"fmt"
	"strings"
)

func stringy(num int) string {
	var collector string

	collector = strings.Repeat("10", (num+1)/2)

	return collector[:num]
}

func main() {
	fmt.Println(stringy(6))
	fmt.Println(stringy(9))
	fmt.Println("vim-go")
}
