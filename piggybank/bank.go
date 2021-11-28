package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	piggyBank := 0.0

	for {
		if piggyBank >= 20.0 {
			fmt.Println("enough money!")
			break
		}

		var deposit float64
		switch rand.Intn(3) {
		case 0:
			deposit = 0.05
		case 1:
			deposit = 0.10
		case 2:
			deposit = 0.25
		}
		piggyBank += deposit
		fmt.Printf("deposit %.2f to piggy bank, now piggy bank is %.2f\n", deposit, piggyBank)
		time.Sleep(time.Second)
	}
}