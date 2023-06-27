package timeComplexity

import "math"

func PrimeNumber(n int) bool {
	if n <= 1 {
		return false
	}

	primes := make([]bool, n+1)

	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if primes[i] {
			for j := i * i; j <= n; j += i {
				primes[j] = false
			}
		}
	}
	return primes[n]
}
