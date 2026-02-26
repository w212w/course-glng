package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	data map[string]int
	mu   sync.Mutex
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		data: make(map[string]int),
	}
}

func (sm *SafeMap) Set(key string, value int) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	sm.data[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	value, exists := sm.data[key]
	return value, exists
}

func main() {
	safeMap := NewSafeMap()

	var wg sync.WaitGroup

	for i := 0; i < 90; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.Set(fmt.Sprintf("key%d", i), i)
		}(i)
	}

	wg.Wait()

	for i := 0; i < 90; i++ {
		if value, exists := safeMap.Get(fmt.Sprintf("key%d", i)); exists {
			fmt.Printf("key%d: %d\n", i, value)
		}
	}
}
