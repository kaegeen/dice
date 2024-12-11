package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rollDice simulates rolling a single six-sided die.
func rollDice() int {
	return rand.Intn(6) + 1
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator

	fmt.Println("Dice Roller")
	fmt.Println("=============")

	// Simulate rolling two dice
	dice1 := rollDice()
	dice2 := rollDice()

	// Display the results
	fmt.Printf("You rolled a %d and a %d.\n", dice1, dice2)
	fmt.Printf("Total: %d\n", dice1+dice2)

	// Optionally, check for special results
	if dice1 == dice2 {
		fmt.Println("You rolled doubles!")
	}
}
