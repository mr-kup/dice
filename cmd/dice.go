package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/kriskiddell/plog"
)

type rollResult struct {
	DiceRolls    []int
	Modifier     int
	ResultString string
	Total        int
}

func main() {
	roll_result, err := rollWithModifier(2, 2, 1)
	if err != nil {
		plog.Error.Println(err)
	}

	fmt.Print(roll_result.ResultString)

}

func rollDie(sides int) (int, error) {

	if sides < 2 || sides > 1000 {
		err := errors.New("Invalid number of sides.")
		return 0, err
	}

	return rand.Intn(sides) + 1, nil

}

func rollDice(number int, sides int) ([]int, error) {

	if number < 1 || number > 1000 {
		return nil, errors.New("Invalid number of dice.")
	}

	if sides < 1 || sides > 1000 {
		return nil, errors.New("Invalid number of sides.")
	}

	results := make([]int, number)

	for i := range number {
		result, err := rollDie(sides)
		if err != nil {
			return nil, err
		}
		results[i] = result
	}

	return results, nil
}

func rollWithModifier(number int, sides int, mod int) (rollResult, error) {
	results, err := rollDice(number, sides)
	if err != nil {
		return rollResult{}, err
	}

	result := 0
	result_string := fmt.Sprintf("%dd%d(", number, sides)
	for i, r := range results {
		result += r
		result_string += strconv.Itoa(r)
		if i < len(results)-1 {
			result_string += " + "
		} else {
			result_string += ")"
		}
	}

	if mod != 0 {
		result += mod
		result_string += fmt.Sprintf("%+d ", mod)
	}

	result_string += fmt.Sprintf("= %d", result)

	rollResult := rollResult{
		DiceRolls:    results,
		Modifier:     mod,
		Total:        result,
		ResultString: result_string,
	}

	return rollResult, nil

}
