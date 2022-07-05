package main

import "fmt"

var i int = 34

var (
	name string = "Ashiq"
	age int = 26
	score float32 = 70.6
)

func main() {
		fmt.Println("i : ",i)

		fmt.Println("Name:",name,"Age:",age,"Score:",score)

		var marks int
		marks = 45
		fmt.Println("marks : ",(1+marks))

		var age int = 26
		fmt.Println("age is : ",age)

		difference := 56.5
		fmt.Println("Difference is : ",difference)

		// value and type
		fmt.Printf("%v is of type %T",age,age)

		var intDifference int
		intDifference = int(difference)

		fmt.Println("\nintDiff : ",intDifference)
}