package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(roll_dice(2000, 20))
}

func roll_die(sides int) int {
	return rand.Intn(sides) + 1
}

func roll_dice(number int, sides int) []int {
	results := make([]int, number)
	for i := range number {
		results[i] = roll_die(sides)
	}
	return results
}
