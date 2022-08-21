package main

import (
	"fmt"
	"strconv"
)

func twice(num int) (double int) {
	str := strconv.Itoa(num)
	half := len(str) / 2

	if str[:half] == str[half:] {
		double = num
	} else {
		double = num * 2
	}

	return
}

func main() {
	fmt.Println(twice(44))
	fmt.Println(twice(444))
	fmt.Println(twice(103103))
	fmt.Println("vim-go")
}
