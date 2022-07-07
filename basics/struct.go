package main

import (
	"fmt"
)

type Person struct {
	name    string
	salary  int
	numbers []int
}

func main() {
	var person Person

	person.name = "ashiq"
	person.salary = 3200
	person.numbers = append([]int{}, []int{1, 2, 3}...)

	fmt.Printf("person %v\n", person)
	fmt.Printf("name %v\n", person.name)
	fmt.Printf("salary %v\n", person.salary)
	fmt.Printf("numbers %v\n", person.numbers)

}
