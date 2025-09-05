package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func workerFlag(stop *bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if *stop {
			fmt.Println("[flag] Worker stopped by flag")
			return
		}
		fmt.Println("[flag] Working...")
		time.Sleep(time.Second)
	}
}

func workerChan(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("[chan] Worker stopped by channel")
			return
		default:
			fmt.Println("[chan] Working...")
			time.Sleep(time.Second)
		}
	}
}

func workerContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[context] Worker stopped by context")
			return
		default:
			fmt.Println("[context] Working...")
			time.Sleep(time.Second)
		}
	}
}

func workerGoexit() {
	for i := 0; i < 5; i++ {
		if i == 2 {
			fmt.Println("[goexit] Worker calls Goexit")
			runtime.Goexit() 
		}
		fmt.Println("[goexit] Working...", i)
		time.Sleep(time.Second)
	}
	fmt.Println("[goexit] This will not be printed")
}

func workerCloseChan(jobs <-chan int) {
	for job := range jobs {
		fmt.Println("[closechan] Got job", job)
		time.Sleep(time.Second)
	}
	fmt.Println("[closechan] Worker stopped by channel close")
}

func workerSignal(stop <-chan struct{}) {
	for {
		select {
		case <-stop:
			fmt.Println("[signal] Worker stopped by OS signal")
			return
		default:
			fmt.Println("[signal] Working...")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	fmt.Println("=== Демонстрация всех способов остановки горутин ===")

	fmt.Println("\n1) Остановка по условию (флаг)")
	var stopFlag bool
	var wg sync.WaitGroup
	wg.Add(1)
	go workerFlag(&stopFlag, &wg)
	time.Sleep(3 * time.Second)
	stopFlag = true
	wg.Wait()

	fmt.Println("\n2) Остановка через канал уведомления")
	stopChan := make(chan struct{})
	go workerChan(stopChan)
	time.Sleep(3 * time.Second)
	close(stopChan)
	time.Sleep(time.Second)

	fmt.Println("\n3) Остановка через context")
	ctx, cancel := context.WithCancel(context.Background())
	go workerContext(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(time.Second)

	fmt.Println("\n4) Остановка через runtime.Goexit()")
	go workerGoexit()
	time.Sleep(5 * time.Second)

	fmt.Println("\n5) Остановка по закрытию рабочего канала")
	jobs := make(chan int)
	go workerCloseChan(jobs)
	for i := 1; i <= 3; i++ {
		jobs <- i
		time.Sleep(time.Second)
	}
	close(jobs)
	time.Sleep(time.Second)
	fmt.Println("\n6) Остановка по SIGINT/SIGTERM")
	stopSignal := make(chan struct{})
	go workerSignal(stopSignal)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	close(stopSignal)
	time.Sleep(time.Second)

}