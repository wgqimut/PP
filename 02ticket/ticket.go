/*
The tickets go to Mars

Start by building a prototype that generates 10 random tickets and displays them in a tabular format with a nice header, as follows:
Spaceline        Days Trip type   Price
=======================================
Virgin Galatic     25 Round-trip  $  96
Virgin Galatic     42 One-way     $  37
SpaceX             34 One-way     $  41
Space Adventures   23 Round-trip  $ 100
Space Adventures   23 One-way     $  50
Virgin Galatic     32 Round-trip  $  84
Virgin Galatic     26 Round-trip  $  94
Space Adventures   29 One-way     $  44
Space Adventures   31 Round-trip  $  86
SpaceX             44 Round-trip  $  72

The table should have four columns:
 The spaceline company providing the service
 The duration in days for the trip to Mars (one-way)
 Whether the price covers a return trip
 The price in millions of dollars
For each ticket, randomly select one of the following spacelines: Space Adventures, SpaceX, or Virgin Galactic.
Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km away from Earth at the time.
Randomly choose the speed the ship will travel, from 16 to 30 km/s.This will determine the duration for the trip
to Mars and also the ticket price. Make faster ships more expen- sive, ranging in price from $36 million to
$50 million. Double the price for round trips.
*/

package main

import (
	"fmt"
	"math/rand"
)

const secondPerDay = 86400

func main() {
	formatStr := "%-16v %4v %-10v %6v\n"
	fmt.Printf(formatStr, "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("=======================================")

	distance := 62100000
	for i := 0; i < 10; i++ {
		var spaceline string
		switch rand.Intn(3) {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galatic"
		}

		speed := 16 + rand.Intn(15)
		days := distance / speed / secondPerDay
		price := 20.0 + speed
		var tripType string
		if rand.Intn(2) == 0 {
			tripType = "One-way"
		} else {
			tripType = "Round-trip"
			price *= 2
		}

		priceStr := fmt.Sprintf("$ %3v", price)
		fmt.Printf(formatStr, spaceline, days, tripType, priceStr)

	}
}
