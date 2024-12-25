package main

import (
	"errors"
	"fmt"
	"github.com/kriskiddell/plog"
	"math/rand"
	"slices"
	"strconv"
)

type rollResult struct {
	DiceRolls    []int
	Modifier     int
	ResultString string
	Total        int
}

func main() {
	roll_result, err := rollWithModifier(3, 6, 0, true, 6)
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

func rollWithModifier(number int, sides int, drop int, highest bool, mod int) (rollResult, error) {
	if drop >= number {
		return rollResult{}, errors.New("Dropping more dice than we are rolling.")
	}

	results, err := rollDice(number, sides)
	if err != nil {
		return rollResult{}, err
	}

	sortedResults := slices.Clone(results) // Clone to avoid mutating the original
	slices.Sort(sortedResults)

	var dropped []int
	if highest {
		dropped = sortedResults[len(sortedResults)-drop:]
	} else {
		dropped = sortedResults[:drop]
	}

	total := 0
	resultString := fmt.Sprintf("Rolling %dd%d%+d", number, sides, mod)

	if drop > 0 {
		resultString += fmt.Sprintf(" drop %s %d", map[bool]string{true: "highest", false: "lowest"}[highest], drop)
	}

	resultString += ":\n ("

	for i, r := range results {
		if slices.Contains(dropped, r) {
			resultString += fmt.Sprintf("!%d", r)                                                  // Mark dropped dice
			dropped = slices.Delete(dropped, slices.Index(dropped, r), slices.Index(dropped, r)+1) // Remove one occurrence
		} else {
			total += r
			resultString += strconv.Itoa(r)
		}
		if i < len(results)-1 {
			resultString += " + "
		}
	}
	resultString += ")"

	if mod != 0 {
		total += mod
		resultString += fmt.Sprintf("%+d", mod)
	}
	resultString += fmt.Sprintf(" = %d", total)

	return rollResult{
		DiceRolls:    results,
		Modifier:     mod,
		Total:        total,
		ResultString: resultString,
	}, nil
}
