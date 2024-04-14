package piscine

func RecursiveFactorial(nb int) int {
	return factorialHelper(nb, 1)
}
func factorialHelper(n, acc int) int {
	// Check for negative numbers and prevent big number which kill memory
	if n > 50 || n < 0 {
		return 0
	}
	// Base case: factorial of 0 is 1
	if n == 0 {
		return acc
	}
	// Recursive case: n * factorial(n-1)
	return factorialHelper(n-1, n*acc)
}
