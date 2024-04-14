package piscine

func Fibonacci(index int) int {
	// Check for negative index
	if index < 0 {
		return -1
	}
	// Base cases: Fibonacci(0) = 0, Fibonacci(1) = 1
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	}
	// Recursive case: Fibonacci(n) = Fibonacci(n-1) + Fibonacci(n-2)
	return Fibonacci(index-1) + Fibonacci(index-2)
}
