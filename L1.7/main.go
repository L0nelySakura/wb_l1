package main

import (
	"fmt"
	"sync"
)

type SaferMap struct {
	mut sync.Mutex
	m map[string]int
}

func (s *SaferMap) Set(key string, value int) {
	s.mut.Lock()
	defer s.mut.Unlock()
	s.m[key] = value
}

func (s *SaferMap) Get(key string) (int, bool) {
	s.mut.Lock()
	defer s.mut.Unlock()
	val, ok := s.m[key]
	return val, ok
}

func NewSaferMap() *SaferMap {
	return &SaferMap{
		m: make(map[string]int, 0),
	}
}

func main() {
	sm := NewSaferMap()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			key := fmt.Sprintf("key-%d", i%10) 
			sm.Set(key, i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 10; i++ {
		key := fmt.Sprintf("key-%d", i)
		val, ok := sm.Get(key)
		if ok {
			fmt.Printf("%s = %d\n", key, val)
		} else {
			fmt.Printf("%s not found\n", key)
		}
	}
}