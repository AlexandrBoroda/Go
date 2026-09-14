package main

import "fmt"

func main() {
	var a, b int = 17, 5

	fmt.Println("a =", a, "b =", b)
	fmt.Println("Сумма        :", a+b)
	fmt.Println("Разность     :", a-b)
	fmt.Println("Произведение :", a*b)
	fmt.Println("Деление      :", a/b)
	fmt.Println("Остаток      :", a%b)
}
