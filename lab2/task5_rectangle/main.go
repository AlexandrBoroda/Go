package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10.5, Height: 4.0}
	fmt.Printf("Прямоугольник: %.1f x %.1f\n", rect.Width, rect.Height)
	fmt.Printf("Площадь: %.2f\n", rect.Area())
}
