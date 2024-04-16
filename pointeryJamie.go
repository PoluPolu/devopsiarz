package main

import "fmt"

func pointeryJamie() {
	fmt.Printf("-----pointery------\n")

	i, j := 42, 6809

	fmt.Printf("i: %d, j: %d\n", i, j)
	fmt.Printf("add i: %v, add j: %v\n", &i, &j)

	p := &i
	fmt.Printf("var p: %v, add i: %v\n",p, &i)

	fmt.Printf("*p: %v\n",*p)
	fmt.Printf("Typ p: %T\n",p)

	*p = 21
	fmt.Printf("i: %d\n", i)

}
