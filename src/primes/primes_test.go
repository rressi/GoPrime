package primes

import (
	"testing"
)

func TestFindPrimes_100(t *testing.T) {

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

}

func TestFindPrimes_10k(t *testing.T) {

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

func TestFindPrimes_1M(t *testing.T) {

	// Next tests requires a reference set of prime numbers:
	findPrimes10k := func () []Number {
		primes10k := make([]Number, 0)
		for _, primeGroup := range FindPrimes(10000) {
			for _, prime := range primeGroup {
				primes10k = append(primes10k, prime)
			}
		}
		return primes10k
	}
	primes_10k := findPrimes10k()


	// Tests with bigger numbers:
	testBig := func (N Number){

		prevPrime := Number(1)
		for _, primeGroup := range FindPrimes(N) {
			for _, prime := range primeGroup {
				for i := 0; i < len(primes_10k) && primes_10k[i] < prime; i++ {
					if prime.isDerivedBy(primes_10k[i]) {
						t.Error(prime, "is not a prime, it can be divided by", primes_10k[i])
					}
				}
				if prime <= prevPrime {
					t.Error(prime, "have been returned after", prevPrime)
				}
				prevPrime = prime
			}
		}
	}

	testBig(100000)
	testBig(1000000)
	testBig(10000000)
	// testBig(100000000)
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
