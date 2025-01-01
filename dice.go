package dice

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type rollResult struct {
	DiceRolls    []int
	Modifier     int
	ResultString string
	Total        int
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

func RollWithModifier(number int, sides int, drop int, highest bool, mod int) (rollResult, error) {

	if drop >= number {
		return rollResult{}, errors.New("Dropping more dice than we are rolling.")
	}

	results, err := rollDice(number, sides)
	if err != nil {
		return rollResult{}, err
	}

	sortedResults := slices.Clone(results)
	slices.Sort(sortedResults)

	var dropped []int
	if highest {
		dropped = sortedResults[len(sortedResults)-drop:]
	} else {
		dropped = sortedResults[:drop]
	}

	resultString := fmt.Sprintf("\n Rolling %dd%d%s", number, sides, map[bool]string{true: fmt.Sprintf("%+d", mod), false: ""}[mod != 0])

	if drop > 0 {
		resultString += fmt.Sprintf(" drop %s %d", map[bool]string{true: "highest", false: "lowest"}[highest], drop)
	}

	total := 0
	resultString += ":\n ("

	for i, r := range results {
		if slices.Contains(dropped, r) {
			resultString += fmt.Sprintf("\u001b[2m%d\u001b[0m", r)                                 // Mark dropped dice
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

func ParseRollString(roll string) (rollResult, error) {
	roll = strings.ReplaceAll(roll, " ", "")
	re := regexp.MustCompile(`(\d+)(d\d+)?(d\d+)?([+-]\d+)?`)
	matches := re.FindStringSubmatch(roll)

	number, err := strconv.Atoi(matches[1])

	if err != nil {
		return rollResult{}, errors.New(fmt.Sprintf("Unable to parse number of dice: %s", err))
	}

	sides, err := strconv.Atoi(strings.Split(matches[2], "d")[1])
	if err != nil {
		return rollResult{}, errors.New(fmt.Sprintf("Unable to parse number of sides: %s", err))
	}

	return RollWithModifier(number, sides, 0, false, 0)

}
