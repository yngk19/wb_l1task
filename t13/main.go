package main

import "fmt"

func main() {
	// создаём две переменные
	a := 5
	b := 10
	// меняем их значения местами
	a, b = b, a
	// печатаем
	fmt.Println(a, b)
}
