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
