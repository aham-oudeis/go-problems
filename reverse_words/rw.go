package main

import (
	"fmt"
	"strings"
)

func reverse(list []rune) []rune {
	l := len(list)
	last := l - 1
	collection := make([]rune, l)

	for i := 0; i < l; i++ {
		collection[last-i] = list[i]
	}

	return collection
}

func reverseWords(str string) string {
	words := strings.Split(str, " ")

	col := make([]string, len(words))

	for i, word := range words {
		if len(word) >= 5 {
			word = string(reverse([]rune(word)))
		}

		col[i] = word
	}

	return strings.Join(col, " ")
}

func main() {
	fmt.Println(reverseWords("Professional"))
	fmt.Println(reverseWords("Walk around the block"))
	fmt.Println("vim-go")
}
