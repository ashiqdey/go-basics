package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)

	go sayHello()

	msg := "hello message"
	go func(msg string) {
		fmt.Println(msg)
		wg.Done()
	}(msg)

	msg = "goodbye"

	wg.Wait()
}

func sayHello() {
	fmt.Println("hello")
	wg.Done()
}
