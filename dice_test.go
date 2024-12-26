package dice

import (
	"testing"
)

func TestRollDie(t *testing.T) {
	sides := []int{2, 4, 6, 8, 10, 12, 20, 69, 100, 420, 888}

	for _, side := range sides {
		highest_roll := 0
		for range 10000 {
			result, _ := rollDie(side)

			if result > highest_roll {
				highest_roll = result
			}

			// Cannot be less than one or greater that the max vaule.
			if result < 1 || result > side {
				t.Errorf("Rolled a %d on a die with %d sides.", result, side)
			}
		}

		// should get at least one max vaule over 10,000 rolls.
		if highest_roll != side {
			t.Errorf("Did not roll at least one %d on a %d sided die, highest was %d", side, side, highest_roll)
		}
	}

}

func TestRollDieErrors(t *testing.T) {

	invaild_sides := []int{-10, -1, 0, 1}

	for _, side := range invaild_sides {
		_, err := rollDie(side)
		if err == nil {
			t.Errorf("RollDie does not error with the invaild number of sides %d", side)
		}
	}
}

func TestRollDice(t *testing.T) {
	number, sides := 2, 20

	result, _ := rollDice(number, sides)

	if len(result) != number {
		t.Errorf("Incorrect number of results.")
	}

}

func TestRollDiceErrors(t *testing.T) {
	number := 1001
	_, err := rollDice(number, 20)
	if err == nil {
		t.Errorf("RollDice does not error with the invaild number of dice %d", number)
	}
}

func TestRollDiceWithModifier(t *testing.T) {
	number, sides, mod := 2, 20, 5

	for range 1000 {
		result, _ := RollWithModifier(number, sides, 0, false, mod)
		if result.Total < number+mod || result.Total > (number*sides)+mod {
			t.Errorf("result out of range")
		}
	}

}
