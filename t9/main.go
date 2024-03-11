package main

import "fmt"

func main() {
	// создаем массив из последовательности чисел
	var array = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// также два канала для реализации конвеера
	firstchan := make(chan int)
	secondchan := make(chan int)
	// горутина пишущая в первый канал
	go func() {
		for _, elm := range array {
			firstchan <- elm
		}
	}()
	// горутина читающая из первого канала и пишущая во второй квадраты чисел
	go func() {
		for range array {
			elm := <-firstchan
			secondchan <- elm * elm
		}
	}()
	// выводим в stdout значения из второго канала
	for range array {
		fmt.Println(<-secondchan)
	}
}
