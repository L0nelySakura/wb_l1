package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)


func worker(id int, jobs <-chan int) {
	for job := range jobs {
		fmt.Printf("\tWorker %d got job %d\n", id, job)
		time.Sleep(time.Second * 10)
		fmt.Printf("\t\tWorker %d completed job %d\n", id, job)
	}
}


func main() {	
	num, _ := strconv.Atoi(os.Args[1])
	
	jobs := make(chan int, 3)
	
	for i := 1; i <= num; i++ {
		go worker(i, jobs)
	}

	current_job := 1
	for {
		jobs <- current_job
		fmt.Printf("Sent job %d\n", current_job)
		current_job++
		time.Sleep(time.Second * 2)
	}

}