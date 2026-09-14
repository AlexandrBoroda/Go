package main

import "fmt"

func main() {
	// Тип выводится автоматически из значения справа
	city := "Москва"       // string
	population := 13000000 // int
	area := 2561.5         // float64
	isCapital := true      // bool

	fmt.Println("Город     :", city)
	fmt.Println("Население :", population)
	fmt.Println("Площадь   :", area, "кв.км")
	fmt.Println("Столица?  :", isCapital)

	// Множественное объявление через :=
	x, y, z := 1, 2, 3
	fmt.Println("x, y, z =", x, y, z)

	fmt.Printf("Типы: city=%T population=%T area=%T isCapital=%T\n",
		city, population, area, isCapital)
}
