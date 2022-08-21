package main

import (
	"fmt"
	"regexp"
)

func cleanUp(str string) string {
	pattern := regexp.MustCompile(`[^a-zA-Z0-9]{1,}`)
	return pattern.ReplaceAllString(str, " ")
}

func main() {
	fmt.Println(cleanUp("---what's my +*& line?"))
}
