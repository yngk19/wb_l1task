package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

// создаем переменную с количеством воркеров. Если пользователь запустит программу с флагом -w, то указанное значение будет в этой переменной
var workersCount = flag.Int("w", 1, "Count of workers")

func main() {
	// парсим флаги
	flag.Parse()
	// объявляем инстанс rand.New
	r := rand.New(rand.NewSource(100))
	// создаём буферизированный канал, с размером workersCount
	mychan := make(chan int, *workersCount)
	// запускаем горутину, котоаря будет писать рандомные числа в канал
	go func() {
		for {
			x := r.Int()
			mychan <- x
		}
	}()
	// запускаем цикл воркеров, которые будут читать из канала и выводить данные в stdout
	for i := 0; i < *workersCount; i++ {
		go func() {
			for {
				o := <-mychan
				fmt.Println(o)
			}
		}()
	}
	// тут создаём канал, который будет принимать объект типа os.Signal - это поможет нам отследить нажатие CTR + C и завершить программу
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT)
	<-signals
}
