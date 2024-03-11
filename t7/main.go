package main

import "sync"

func main() {
	// создаем объект мьютекса для синхронизации
	var mtx sync.Mutex
	// мапу для конкурентной записи
	mymap := make(map[int]int)
	// используем паттерн WorkerPool
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			// регистрируем завершение воркера после выполнения горутины
			defer wg.Done()
			// блокируем мапу для записи
			mtx.Lock()
			mymap[i] = i
			// разблокировка
			mtx.Unlock()
		}()
	}
	// ожидаем завершения воркеров
	wg.Wait()
}
