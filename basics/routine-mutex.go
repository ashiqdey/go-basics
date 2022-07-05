package main

import (
	"fmt"
	// "runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter int = 0
var m = sync.RWMutex{}

func main() {
	// runtime.GOMAXPROCS(10)
	for i := 0; i < 5; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increament()
	}

	wg.Wait()
}

func sayHello() {
	fmt.Printf("#count %v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increament() {
	counter++
	m.Unlock()
	wg.Done()
}
