package main

import "fmt"

func leadingSubstrings(str string) []string {
	col := []string{}

	for idx := range str {
		col = append(col, str[:idx+1])
	}

	return col
}

func allSubstrings(str string) []string {
	col := []string{}

	for idx := range str {
		col = append(col, leadingSubstrings(str[idx:])...)
	}

	return col
}

func main() {
	fmt.Println(leadingSubstrings("abc"))
	fmt.Println(allSubstrings("abcde"))
	fmt.Println("vim-go")
}
