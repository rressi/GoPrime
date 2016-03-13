package primes

import (
	"testing"
)

func TestFindPrimes(t *testing.T) {

	// Tests range 0 - 100
	for n := Number(0); n <= 100; n++ {

		primes := FindPrimes(n)
		// t.Log(n, "->", primes)

		i := 0
		expectedPrime := PRIMES_100[i]
	checkLoop:
		for _, primeGroup := range primes {
			for _, prime := range primeGroup {
				if prime != expectedPrime {
					t.Error("Value", prime, "returned while", expectedPrime, "was expected")
					break checkLoop
				}
				i++
				expectedPrime = PRIMES_100[i]
			}
		}
	}

	// Tests range [101, 10000]
	for n := Number(101); n <= 10000; n++ {
		primes := FindPrimes(n)
		// t.Log(n, "->", primes)
		for _, primeGroup := range primes {
			for _, prime := range primeGroup {
				for i := 0; i < len(PRIMES_100) && PRIMES_100[i] < prime; i++ {
					if prime.isDerivedBy(PRIMES_100[i]) {
						t.Error(prime, "is not a prime, it can be divided by", PRIMES_100[i])
					}
				}
			}
		}
	}
}

func BenchmarkFindPrimes_10_9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindPrimes(1000000000)
	}
}

func BenchmarkFindPrimes_10_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindPrimes(10000000000)
	}
}
