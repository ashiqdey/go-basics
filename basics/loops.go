package main

import (
	"fmt"
)


func main() {

	for i := 0; i<5; i++ {
		fmt.Println(i)
	}

		languages := map[string]int{
		"python" : 145,
		"golang" : 98,
		"java" : 67,
	}

	for k,v := range languages {
		fmt.Println(k,v)
	}

}