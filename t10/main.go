package main

import "fmt"

func main() {
	// cоздаем массив из температур
	temps := [...]float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	// также слайс для их хранения
	answer := make(map[int][]float32)
	for _, elm := range temps {
		// заносим значения по ключам с шагом в 10 градусов
		k := int(elm/10) * 10
		answer[k] = append(answer[k], elm)
	}
	// печатаем значения
	for key, elm := range answer {
		fmt.Println(key, elm)
	}
}
