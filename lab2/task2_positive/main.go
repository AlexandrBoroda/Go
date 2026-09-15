package main

import "fmt"

func signOfNumber(n int) string {

	var num int

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	if n > 0 {
		return "Positive"
	} else if n < 0 {
		return "Negative"
	} else {
		return "Zero"
	}
}

func main() {
	fmt.Println(signOfNumber(5))  // Positive
	fmt.Println(signOfNumber(-3)) // Negative
	fmt.Println(signOfNumber(0))  // Zero
}
