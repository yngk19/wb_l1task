package main

import (
	"fmt"
	"sync"
)

// функция для вычисления квадрата и печати в stdout
func PrintSquare(wg *sync.WaitGroup, number int) {
	// регистрируем wg.Done() для уменьшения счетчика горутин
	defer wg.Done()

	fmt.Println(number * number)
}

func main() {
	array := []int{2, 4, 6, 8, 10}
	// обьявляем переменную типа sync.WaitGroup для реализации паттерна пул воркеров (Worker pool)
	var wg sync.WaitGroup
	for _, elm := range array {
		// добавляем в пулл 1 воркер
		wg.Add(1)
		go PrintSquare(&wg, elm)
	}
	// ждем завершения выполнения всех воркеров
	wg.Wait()
}
