// Write a function that returns the first prime number that is equal or superior to the int passed as parameter.
// The function must be optimized in order to avoid time-outs with the tester.
// (We consider that only positive numbers can be prime numbers)

package piscine

func FindNextPrime(n int) int {
	// If n is less than or equal to 1, return 2 (the first prime number)
	if n <= 1 {
		return 2
	}
	// Start checking numbers from n onwards until a prime number is found
	for i := n; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
}
