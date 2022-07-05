package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(10)
	fmt.Printf("thread count %v\n", runtime.GOMAXPROCS(-1))
}
