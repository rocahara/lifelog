package main

import (
	"fmt"
	"sync"
	"time"
)

type MutexTest struct {
	mu sync.Mutex
	n  int
}

func main() {
	fmt.Println("test")
	wg := sync.WaitGroup{}

	m := &MutexTest{
		n: 3,
	}

	wg.Add(1)
	go func() {
		m.Add(&wg)
	}()

	time.Sleep(100 * time.Millisecond)
	wg.Add(1)
	go func() {
		m.Read(&wg)
	}()

	wg.Wait()
}

func (m *MutexTest) Add(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Add start")
	m.mu.Lock()
	time.Sleep(5 * time.Second)
	m.n++
	m.mu.Unlock()
	fmt.Println("Add end")
}

func (m *MutexTest) Read(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Read start")
	m.mu.Lock()
	fmt.Printf("Read:%d\n", m.n)
	m.mu.Unlock()
	fmt.Println("Read end")
}
