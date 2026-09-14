package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("RFC1123 :", now.Format(time.RFC1123))
	fmt.Println("RFC3339 :", now.Format(time.RFC3339))
	fmt.Println("Дата    :", now.Format("02.01.2006"))
	fmt.Println("Время   :", now.Format("15:04:05"))
	fmt.Println("Полный  :", now.Format("02.01.2006 15:04:05 Monday"))
}
