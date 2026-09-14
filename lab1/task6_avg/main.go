package main

import "fmt"

// average — среднее трёх чисел
func average(a, b, c float64) float64 {
	return (a + b + c) / 3
}

func main() {
	// Вариант 1: сразу float64
	x, y, z := 3.5, 7.0, 10.5
	fmt.Printf("Среднее (%.1f, %.1f, %.1f) = %.4f\n",
		x, y, z, average(x, y, z))

	// Вариант 2: целые на входе, дробный результат
	a, b, c := 10, 20, 33
	avg := float64(a+b+c) / 3 // приведение ДО деления
	fmt.Printf("Среднее (%d, %d, %d) = %.4f\n", a, b, c, avg)
}
