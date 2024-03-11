package main

import "fmt"

func main() {
	// создаём два пустых множества
	set1 := make(map[int]struct{})
	set2 := make(map[int]struct{})
	// создаем множество в котором будет хранится пересечение предыдущих двух множеств
	intersectionset := make(map[int]struct{})
	// заполняем те два множества значениями; будут также повторяющиеся значения
	for i := 0; i < 10; i++ {
		set1[i] = struct{}{}
		if i > 4 {
			set2[i] = struct{}{}
		}
	}
	// находим пересечение множеств и добавляем эти элементы в третье множество
	for key, _ := range set1 {
		_, ok := set2[key]
		if ok {
			intersectionset[key] = struct{}{}
		}
	}
	// печатаем элементы пересечения множеств
	for key, _ := range intersectionset {
		fmt.Println(key)
	}
}
