package main

import (
	"fmt"
)

func main() {

	defer cleanup()

	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i == 2 {
			panic("Number cant be 2")
		}
	}

}

func cleanup() {
	if err := recover(); err != nil {
		fmt.Println("Recoved", err)
	}
}
