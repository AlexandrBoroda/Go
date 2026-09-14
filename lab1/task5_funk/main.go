package main

import "fmt"

// sumAndDiff возвращает два значения сразу.
// В Go это стандартная практика — не нужны структуры или указатели.
func sumAndDiff(x, y float64) (float64, float64) {
	return x + y, x - y
}

func main() {
	var a, b float64 = 10.5, 4.25

	// Принимаем два возвращаемых значения
	sum, diff := sumAndDiff(a, b)

	fmt.Printf("a = %.2f, b = %.2f\n", a, b)
	fmt.Printf("Сумма    = %.2f\n", sum)
	fmt.Printf("Разность = %.2f\n", diff)
}
