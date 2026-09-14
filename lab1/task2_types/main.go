package main

import "fmt"

func main() {
	// Полная форма объявления: var имя тип = значение
	var age int = 25
	var pi float64 = 3.14159
	var name string = "Иван"
	var isStudent bool = true

	fmt.Println("age       =", age)
	fmt.Println("pi        =", pi)
	fmt.Println("name      =", name)
	fmt.Println("isStudent =", isStudent)

	// %T в Printf показывает тип переменной
	fmt.Printf("Типы: %T, %T, %T, %T\n", age, pi, name, isStudent)
}
