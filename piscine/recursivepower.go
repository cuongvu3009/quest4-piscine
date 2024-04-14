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
