package main

import "fmt"


func main() {
	// boolean
		var admin bool = true
		fmt.Printf("admin %v, type %T\n",admin,admin)

		old := 1==(2-1)
 		fmt.Printf("old %v, type %T\n",old,old)

		// numeric
		n :=42
		fmt.Printf("n %v, type %T\n",n,n)

		// var avg float32 = 35.6
		avg := 3.4
		// avg  = 13.4E3
		// avg  = 2.1E14 
		fmt.Printf("avg %v, type %T\n",avg,avg)

		a := 10.2
		b := 3.7
		fmt.Println("sum ",(a+b))

		// complex number
		complex := 1 + 2i
		fmt.Printf("complex %v, type %T\n",complex,complex)
		fmt.Printf("real %v, imag %v\n",real(complex),imag(complex))

		// string utf8
		name := "Ashiq"
		fmt.Printf("%v , %T\n",name,name)
		fmt.Printf("%v , %T",string(name[2]),name[2])

		// convert to slice of byte
		bytes := []byte(name)
		fmt.Printf("%v , %T\n",bytes,bytes)


		// rune utf32
		// var r rune = 'a'
		r := 'a' // always use '
		fmt.Printf("%v , %T\n",r,r)

}