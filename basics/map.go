package main

import (
	"fmt"
)

func main() {
	/*
			var a = make(map[string]string)
		  a["brand"] = "Ford"
		  a["model"] = "Mustang"
		  a["year"] = "1964"
	*/

	languages := map[string]int{
		"python": 145,
		"golang": 98,
		"java":   67,
	}

	fmt.Printf("languages %v", languages)

}
