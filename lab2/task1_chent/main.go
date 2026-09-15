package main

import "fmt"

func main() {
	var num int

	fmt.Print("Введите число: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println(num, "— чётное число")
	} else {
		fmt.Println(num, "— нечётное число")
	}
}
