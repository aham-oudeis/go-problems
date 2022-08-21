package main

import (
	"fmt"
	"strings"
)

func wordCap(str string) string {
	return strings.Title(str)
}

func main() {
	fmt.Println(wordCap("four score and seven"))
	fmt.Println(strings.Title("\"hi\""))
	fmt.Println(wordCap("this is a \"quoted\" word"))
	fmt.Println("vim-go")
}
