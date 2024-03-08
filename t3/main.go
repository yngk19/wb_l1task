package main

import (
	"fmt"
	"sync"
)

func main() {
	// создаем слайс с элементами последовательности
	var seq = []int{2, 4, 6, 8, 10}
	// объявляем объект sync.WaitGroup
	var wg sync.WaitGroup
	// добавляем 5 воркеров
	wg.Add(5)
	for _, elm := range seq {
		go func(elm int) {
			// регистрируем завершение воркера
			defer wg.Done()
			fmt.Println(elm * elm)
		}(elm)
	}
	// ждем завершения всех воркеров
	wg.Wait()
}
