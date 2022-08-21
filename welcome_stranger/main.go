package main

import (
	"fmt"
	"strings"
)

type Person struct {
	name       []string
	title      string
	occupation string
}

func main() {
	name := []string{"John", "Q", "Doe"}
	john := Person{name, "Master", "Plumber"}

	fullname := strings.Join(name, " ")

	message := fmt.Sprintf("Hello, %s! Nice to have a %s %s around", fullname, john.title, john.occupation)
	fmt.Println(message)
}
