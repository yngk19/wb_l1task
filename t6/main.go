package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Все возможные способы остановки горутины

	// 1-й - принять сигнал от ОС
	sygnals := make(chan os.Signal)
	signal.Notify(sygnals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	<-sygnals
	// 2-й - сделать siscall exit с кодом 1
	os.Exit(1)
	// 3-й - сделать break в цикле при выполнении условия
	condition := 1
	go func() {
		for {
			// do something
			if condition == 1 {
				break
			}
		}
	}()
	// 4-й - sync.WaitGroup паттерн WorkerPool
	var wg *sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		// do something
		wg.Done()
	}(wg)
	wg.Wait()
	// 5-й - select {}
	mychan := make(chan bool)
	go func() {
		for {
			select {
			case <-mychan:
				break
			default:
				// do something
			}
		}
	}()
	// 6-й - context
	go func(ctx context.Context) {
		for {
			// do something
			select {
			case <-ctx.Done():
				break
			}
		}
	}(ctx)
}
