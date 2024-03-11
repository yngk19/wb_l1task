package main

import (
	"flag"
	"fmt"
	"time"
)

// время работы программы
var N = flag.Int("N", 10, "Time to live a program")

func main() {
	flag.Parse()
	mychan := make(chan int)
	// горутина, пишущая в канал
	go func() {
		for {
			mychan <- 100
		}
	}()
	// горутина читающая из канала
	go func() {
		for {
			fmt.Println(<-mychan)
		}
	}()
	// ожидаем N секунд
	time.Sleep(time.Duration(*N) * time.Second)
}
