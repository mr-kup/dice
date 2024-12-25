package main

import "testing"

func TestRollDie(t *testing.T) {
	sides := []int{2, 4, 6, 8, 10, 12, 20, 69, 100, 420, 888}

	for _, side := range sides {
		highest_roll := 0
		for range 10000 {
			result := rollDie(side)

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
