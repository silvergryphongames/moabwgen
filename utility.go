package main

import "math/rand"

func rollDice(number int, sides int) int {
	total := 0
	result := 0

	for i := 0; i < number; i++ {
		result = rand.Intn(sides) + 1
		total += result
	}

	return total
}
