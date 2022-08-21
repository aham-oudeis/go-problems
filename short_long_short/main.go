package main

import "fmt"

func shortLongShort(s1, s2 string) string {
	if len(s1) > len(s2) {
		return shortLongShort(s2, s1)
	}

	return s1 + s2 + s1
}

func main() {
	fmt.Println(shortLongShort("abc", "defgfh"))
	fmt.Println(shortLongShort("abc", "dh"))
	fmt.Println(shortLongShort("", "defgfh"))
}
