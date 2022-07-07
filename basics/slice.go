package main

import (
	"fmt"
)

func main() {
	// array
	// marks := [3]int{65,78,99}
	a := []int{1, 2, 3, 4, 5, 6}
	b := a
	b[1] = 5

	fmt.Printf("slice a %v, type %T, Length %v, Cap %v, \n", a, a, len(a), cap(a))
	fmt.Printf("slice b %v,\n", b)

	// [inclusice, exclusive]
	c := a[:]
	d := a[1:]
	e := a[:2]
	f := a[1:2]

	fmt.Printf("slice c %v,\n", c)
	fmt.Printf("slice d %v,\n", d)
	fmt.Printf("slice e %v,\n", e)
	fmt.Printf("slice f %v,\n", f)

	g := make([]int, 3, 5)
	g = append(g, 2)
	fmt.Printf("slice g %v,\n", g)

	h := []int{}
	h = append(g, 2)
	fmt.Printf("h %v, len %v, cap %v\n", h, len(h), cap(h))

	h = append([]int{}, []int{3, 4, 5, 6}...)
	fmt.Printf("h modified %v, len %v, cap %v\n", h, len(h), cap(h))

}
