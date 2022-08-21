package main

import "fmt"

func isBalanced(str string) bool {
	count := 0

	for i := range str {
		char := string(str[i])

		if char == "(" {
			count += 1
		}

		if char == ")" {
			count -= 1
		}

		if count < 0 {
			return false
		}
	}

	return count == 0
}

func main() {
	fmt.Println(isBalanced("What (is)) this?"))
	fmt.Println("vim-go")
}
