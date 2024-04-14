// Write a recursive function that returns the value of nb to the power of power.
// Negative powers will return 0. Overflows do not have to be dealt with.
// for is forbidden for this exercise.

package piscine

func RecursivePower(nb int, power int) int {
	// If power is negative, return 0
	if power < 0 {
		return 0
	}
	// Base case: when power is 0, return 1
	if power == 0 {
		return 1
	}
	// Recursive case: nb * nb^(power-1)
	return nb * RecursivePower(nb, power-1)
}
