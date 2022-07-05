package main

import "fmt"


func main() {
	// array
	// marks := [3]int{65,78,99}
	marks := [...]int{65,78,99}
	
		fmt.Printf("marks %v, type %T\n",marks,marks)
		fmt.Printf("%v\n",cap(marks))

	var students [3]string
		students[0] = "ashiq"
		students[1] = "jhon"
		students[2] = "joe"
		fmt.Printf("Students %v\n",students)
		fmt.Printf("#1 %v\n",students[1])
		

		//copy array
		a := [...]int{1,2,3}
		b := a
		b[1] = 5
		fmt.Printf("a %v\n",a)
		fmt.Printf("b %v\n",b)

		//copy address
		c := &a
		c[1] = 5
		fmt.Printf("a %v\n",a)
		fmt.Printf("c %v\n",c)

		d := a[:]
		d[1] = 7
		fmt.Printf("a %v\n",a)
		fmt.Printf("d %v\n",d)


}