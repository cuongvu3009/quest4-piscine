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
