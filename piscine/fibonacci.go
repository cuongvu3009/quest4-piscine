// Write a recursive function that returns the value at the position index in the fibonacci sequence.
// The first value is at index 0.
// The sequence starts this way: 0, 1, 1, 2, 3 etc...
// A negative index will return -1.
// for is forbidden for this exercise.

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
