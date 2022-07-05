package main

import (
	"fmt"
)


func main() {


	if 1==1 {
		fmt.Println("#1 true")
	}

		languages := map[string]int{
		"python" : 145,
		"golang" : 98,
		"java" : 67,
	}

	if lang,ok := languages["java"]; ok{
		fmt.Printf("#1 %v \n",lang)

	}

}