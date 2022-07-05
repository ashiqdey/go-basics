package main

import "fmt"

func main() {
	sayMessage("hello")

	//pass pointer
	name := "ashiq"
	changeName(&name)
	fmt.Printf("Name : %v", name)

	fmt.Printf("\nSum is %v\n", sum(1, 2, 3, 4, 5, 6))

	d, err := divide(5.6, 1.1)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Divide : %v\n", d)

	func(name string) {
		fmt.Printf("Name is %v\n", name)
	}("ashiq")

}

func sayMessage(mgs string) {
	fmt.Println(mgs)
}

func changeName(name *string) {
	*name = "ashiq v2"
}

func sum(v ...int) int {
	result := 0
	for _, v := range v {
		result += v
	}

	return result
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}

	return (a / b), nil
}
