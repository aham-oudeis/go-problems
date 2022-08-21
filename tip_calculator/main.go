package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getNumInput(msg string) (float64, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(msg)
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)

	num, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, false
	}

	return num, true
}

func main() {
	total, _ := getNumInput("What's the total bill? ")
	tipPercent, _ := getNumInput("What's the tip percentage? ")

	fmt.Println("the tip amounts to $", (total/100)*tipPercent)
}
