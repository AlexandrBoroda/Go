package main

import (
	"fmt"
	"unicode/utf8"
)

func stringLength(s string) int {
	return utf8.RuneCountInString(s)
}

func main() {
	fmt.Println(stringLength("Hello"))  // 5
	fmt.Println(stringLength("Привет")) // 6
	fmt.Println(stringLength(""))       // 0
}
