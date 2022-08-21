package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter the length: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(input)

	input = strings.TrimSpace(input)
	fmt.Printf("the type of input is %T", input)
	fmt.Println(strconv.ParseFloat(input, 64))
	length, _ := strconv.ParseFloat(input, 64)
	fmt.Println("the length of the room is: ", length)
}
