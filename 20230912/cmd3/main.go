package main

import (
	"fmt"
	"sync"
	"time"
)

type MutexTest struct {
	mu sync.RWMutex
	n  int
}

func main() {
	wg := sync.WaitGroup{}

	m := &MutexTest{
		n: 3,
	}

	wg.Add(1)
	go func() {
		m.Add("add1", &wg)
	}()

	time.Sleep(100 * time.Millisecond)
	wg.Add(1)
	go func() {
		m.Add("add2", &wg)
	}()

	time.Sleep(100 * time.Millisecond)
	wg.Add(1)
	go func() {
		m.Read("read1", &wg)
	}()

	time.Sleep(100 * time.Millisecond)
	wg.Add(1)
	go func() {
		m.Read("read2", &wg)
	}()

	wg.Wait()
	fmt.Printf("m.n: %v\n", m.n)
}

func (m *MutexTest) Add(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(name + " start")
	m.mu.Lock()
	time.Sleep(5 * time.Second)
	m.n++
	m.mu.Unlock()
	fmt.Println(name + " end")
}

func (m *MutexTest) Read(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(name + " start")
	m.mu.RLock()
	fmt.Printf(name+": %d\n", m.n)
	m.mu.RUnlock()
	fmt.Println(name + " end")
}
