// Write a function that returns true if the int passed as parameter is a prime number. Otherwise it returns false.
// The function must be optimized in order to avoid time-outs with the tester.
// (We consider that only positive numbers can be prime numbers)
// (We also consider that 1 is not a prime number)

package piscine

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	// Prime numbers are of the form 6k ± 1, where k is an integer.
	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}
