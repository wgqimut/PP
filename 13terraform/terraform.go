/*

Write a program to terraform a slice of strings by prepending
each planet with "New ". Use your program to terraform Mars,
nus, and Neptune.
Your first iteration can use a terraform function, but your
final implementation should introduce a Planets type with
a terraform method
*/
package main

import "fmt"

type Planets []string

func (p Planets) terraform() {
	for i := range p {
		p[i] = "New " + p[i]
	}
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	Planets(planets[3:5]).terraform()

	fmt.Println(planets)
}
