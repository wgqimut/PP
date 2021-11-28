/*
Write a program that uses a loop to continuously append
an element to a slice. Print out the capacity of the
slice whenever it changes. Does append always double
the capacity when the underlying array runs out of room?
*/
package main

import "fmt"

func main() {
	a := []int{}
	lastCap := cap(a)

	for i := 0; i < 10000; i++ {
		a = append(a, i)
		if cap(a) != lastCap {
			fmt.Printf("slice cap = %d\n", cap(a))
			lastCap = cap(a)
		}
	}
}
