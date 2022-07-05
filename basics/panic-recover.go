package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Start")
	paniker()
	fmt.Println("end")

}

func paniker() {
	fmt.Println("paniker start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("ERROR:", err)
		}
	}()
	panic("Opps! panicking")
	fmt.Println("paniker end")

}
