package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)


func worker(ctx context.Context, id int, jobs <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\tWorker %d shutting down...", id)
			return
		
		case job := <-jobs:
			fmt.Printf("\tWorker %d got job %d\n", id, job)
			time.Sleep(time.Second * 10)
			fmt.Printf("\t\tWorker %d completed job %d\n", id, job) 
		}
	}
}


func main() {	
	num, _ := strconv.Atoi(os.Args[1])
	
	jobs := make(chan int, 3)
	
	ctx, cancel := context.WithCancel(context.Background())


	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	

	for i := 1; i <= num; i++ {
		go worker(ctx, i, jobs)
	}

	go func() {
		<-sig
		fmt.Printf("Received interruption...\n")
		cancel()
		close(jobs)
	}()




	current_job := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Shutting down...\n")
			return

		default: 
			jobs <- current_job
			fmt.Printf("Sent job %d\n", current_job)
			current_job++
			time.Sleep(time.Second * 2)
		}
	}

}