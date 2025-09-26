package main

import (
	"fmt"
	"sync"
)

type Counter struct{
	mu 	sync.Mutex
	val int
}


func newCounter() *Counter {
	return &Counter{}
}


func (c *Counter) Increm() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *Counter) Val() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.val
}


func main() {
	count := newCounter()
	
    var wg sync.WaitGroup

	routines := 10
	increm := 100

	wg.Add(routines)

	for i := 0; i < routines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < increm; j++ {
				count.Increm()
			}
		}(i)
	}

	
    fmt.Printf("Ожидаемое: %d\n", routines * increm)
    fmt.Printf("Фактическое: %d\n", count.Val())
}