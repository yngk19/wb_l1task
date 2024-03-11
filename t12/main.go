package main

import "fmt"

func main() {
	// создаём массив из повторяющихся строк
	wordsSeq := [...]string{"cat", "dog", "dog", "word", "cat", "word", "dog"}
	// создаём пустое множество
	set := make(map[string]struct{})
	// заполняем это множество элементами массива
	for _, elm := range wordsSeq {
		set[elm] = struct{}{}
	}
	// выводим в stdout элементы множества
	for key, _ := range set {
		fmt.Println(key)
	}
}
