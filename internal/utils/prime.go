package utils

import (
	"errors"
)

// SieveOfEratosthenes to generate all prime
// numbers less than n
func SieveOfEratosthenes(n int, isPrime []bool) {
	// Initialize all entries of boolean array
	// as true. A value in isPrime[i] will finally
	// be false if i is Not a prime, else true
	// bool isPrime[n+1];
	for i := 2; i <= n; i++ {
		isPrime[i] = true

		for p := 2; p*p <= n; p++ {
			// If isPrime[p] is not changed, then it is
			// a prime
			if isPrime[p] == true {
				// Update all multiples of p
				for j := p * 2; j <= n; j += p {
					isPrime[j] = false
				}
			}
		}
	}
}

// FindPrimePair to print a prime pair
// with given product
func FindPrimePair(n int) (int, int, error) {
	// Generating primes using Sieve
	isPrime := make([]bool, n+1)
	SieveOfEratosthenes(n, isPrime)

	// Traversing all numbers to find first
	// pair
	for i := 2; i < n; i++ {
		x := n / i

		if isPrime[i] && isPrime[x] && x != i && x*i == n {
			return i, x, nil
		}
	}

	return 0, 0, errors.New("not found")
}
