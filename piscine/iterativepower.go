// Write an iterative function that returns the value of nb to the power of power.
// Negative powers will return 0. Overflows do not have to be dealt with.

package piscine

func IterativePower(nb int, power int) int {
	// If power is negative, return 0
	if power < 0 {
		return 0
	}
	result := 1
	// Multiply nb by itself power times
	for i := 0; i < power; i++ {
		result *= nb
	}
	return result
}
