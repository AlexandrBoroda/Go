package main

import "fmt"

func average(a, b int) float64 {
	return float64(a+b) / 2.0
}

func main() {
	fmt.Printf("Среднее (3, 5) = %.2f\n", average(3, 5))
	fmt.Printf("Среднее (10, 20) = %.2f\n", average(10, 20))
	fmt.Printf("Среднее (7, 8) = %.2f\n", average(7, 8))
}
