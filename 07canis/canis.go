package main

import "fmt"

func main() {
	const lightSpeed = 299792
	const secondPerDay = 86400
	const distance = 236000000000000000
	const daysPerYear = 365

	lightYear := distance / lightSpeed / secondPerDay / daysPerYear
	fmt.Println("the distance is", lightYear, "light yeas away")
}

