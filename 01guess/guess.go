/*
The guess game

Write a guess-the-number program. Make the computer pick random
numbers between 1–100 until it guesses your number, which you
declare at the top of the program.
Display each guess and whether it was too big or too small
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := 42

	min := 1
	max := 100
	rand.Seed(time.Now().Unix())
	for {
		guess := rand.Intn(max-min+1) + min
		if guess > num {
			fmt.Printf("%v is too big!\n", guess)
			max = guess - 1
		} else if guess < num {
			fmt.Printf("%v is too small!\n", guess)
			min = guess + 1
		} else {
			fmt.Printf("You got it!\n")
			break
		}
		time.Sleep(time.Second)
	}
}
