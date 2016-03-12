package primes

import (
	"reflect"
	"testing"
)

func TestFindPrimes(t *testing.T) {

	PRIMES_100 := []Number{
		2,	3,	5,	7,	11,	13,	17,	19,	23,
		29,	31,	37,	41,	43,	47,	53,	59,	61,	67,
		71,	73,	79,	83,	89,	97,
	}

	// Tests range 0 - 100
	for n := Number(0); n <= 100; n++ {
		numExpected := 0
		for numExpected < len(PRIMES_100) && PRIMES_100[numExpected] < n {
			numExpected++
		}
		expectedPrimes := PRIMES_100[:numExpected]

		// t.Log("Testing", n)
		primes := FindPrimes(n)

		if !reflect.DeepEqual(primes, expectedPrimes) {
			t.Errorf("Unexpected values %v != %v", primes, expectedPrimes)
		}
	}

	// Tests range 101 - 10000
	for n := Number(101); n <= 10000; n++ {
		// t.Log("Testing", n)
		primes := FindPrimes(n)
		for _, prime := range primes {
			for i := 0; i < len(PRIMES_100) && PRIMES_100[i] < prime; i++ {
				if prime % PRIMES_100[i] == 0 {
					t.Error(prime, "is not a prime, it can be divided by", PRIMES_100[i])
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
