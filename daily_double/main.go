package main

import "fmt"

func crunch(str string) string {
	var collection string
	var last string

	for i := 0; i < len(str); i++ {
		char := string(str[i])
		if char != last {
			collection = collection + char
			last = char
		}
	}

	return collection
}

func main() {
	crunched := crunch("ddaiilllyy ddooubbllee")
	fmt.Println(crunched)
}
