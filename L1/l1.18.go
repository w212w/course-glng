package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type AtomicCounter struct {
	value int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

type MutexCounter struct {
	mu    sync.Mutex
	value int64
}

func (m *MutexCounter) Inc() {
	m.mu.Lock()
	m.value++
	m.mu.Unlock()
}

func (m *MutexCounter) Value() int64 {
	m.mu.Lock()
	v := m.value
	m.mu.Unlock()
	return v
}

func main() {
	var c AtomicCounter
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup
	var m MutexCounter

	goroutines := 100
	intoGoroutine := 300

	startAtomic := time.Now()
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < intoGoroutine; j++ {
				c.Inc()
			}
		}()
	}
	wg.Wait()
	atomicDur := time.Since(startAtomic)

	startMutex := time.Now()
	wg2.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg2.Done()
			for j := 0; j < intoGoroutine; j++ {
				m.Inc()
			}
		}()
	}
	wg2.Wait()

	mutexDur := time.Since(startMutex)

	fmt.Println("final counter (atomic):", c.Value(), "/ time:", atomicDur)
	fmt.Println("final counter (mutex):", m.Value(), "/ time:", mutexDur)

}
