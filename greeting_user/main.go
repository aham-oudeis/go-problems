package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getInput() string {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(input)
}

func main() {
	fmt.Print("What is your name ? ")
	name := getInput()

	message := fmt.Sprintf("Hello %s.", name)

	if strings.HasSuffix(name, "!") {
		name = name[:len(name)-1]
		message = "HELLO " + strings.ToUpper(name) + ". WHY ARE WE SCREAMING?"
	}

	fmt.Println(message)
}
