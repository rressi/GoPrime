package primes

import (
	"sort"
)

// A natural number
type Number uint64

var PRIMES_100 []Number = []Number {
	2, 3, 5, 7, 11, 13, 17, 19, 23,
	29, 31, 37, 41, 43, 47, 53, 59, 61, 67,
	71, 73, 79, 83, 89, 97, 101}

// Finds prime numbers below N
// Returns primes split in blocks for its convenience, they are sorted
func FindPrimes(N Number) (results [][]Number) {

	// fmt.Println("findNonPrimes(", N, nonPrimes, ")")
	// defer func() {
	// 	fmt.Println("    nonPrimes", nonPrimes)
	// }()

	// Primes below 100 are hard coded:
	if N <= 100 {
		var i int
		for PRIMES_100[i] < N { i++ }
		results = [][]Number{PRIMES_100[:i]}
		return
	}

	// Recursively finds non-primes below or equal to sqrt(N)
	n := (N.sqrt() + 1)
	results = FindPrimes(n)

	if n < 1000 {

		// Fills serially the buffer of non primes:
		primes := findPrimesInRange(n, N, results)
		if len(primes) > 0 {
			results = append(results, primes)
		}

	} else {
		// Fills on parallel the buffer of non primes:
		numResults := Number(len(results))
		newResults := make(chan []Number)
		numRoutines := 0
		for x := n; x < N; x += n {
			go func(x Number) {
				newResults <- findPrimesInRange(x, (x + n).min(N), results)
			}(x)
			numRoutines++
		}
		for i := 0; i < numRoutines; i++ {
			primes := <-newResults
			if len(primes) > 0 {
				results = append(results, primes)
			}
		}

		// Sorts only new results:
		sort.Sort(ByFirst(results[numResults:]))
	}

	return
}

// Finds prime numbers in the range [x0, x1) given all primes that are <= sqrt(x0)
func findPrimesInRange(x0, x1 Number, primes [][]Number) (newPrimes []Number) {

	// fmt.Println("findPrimesFrom", x0, x1, primes)
	// defer func() {
	//   	fmt.Println("    newPrimes", newPrimes)
	// }()

	flags := make([]bool, x1-x0)
	// numPrimes := len(flags)

	for _, primeGroup := range primes {
		for _, prime := range primeGroup {
			for i := int(x0.alignUp(prime) - x0); i < len(flags); i += int(prime) {
				flags[i] = true
				/*
				if !flags[i] {
					flags[i] = true
					numPrimes--
				}
				*/
			}
		}
	}
	// fmt.Println("    flags", flags)

	// Collects new prime numbers:
	estimatedNumPrimes := (x1 - x0) / (x1-x0).log10().max(1)
	newPrimes = make([]Number, 0, estimatedNumPrimes) // , numPrimes)
	for i, hasDividends := range flags {
		if !hasDividends {
			newPrimes = append(newPrimes, x0 + Number(i))
		}
	}

	return
}

// Returns the squared root of a number strictly positive
func (x Number) sqrt() (y Number) {

	for y = 1; y*y <= x; y++ {
	}
	y--

	return
}

// Aligns x to the closest multiple of k that is less or equal to x
func (x Number) align(k Number) (y Number) {
	y = x - (x % k)
	return
}

// Aligns x to the closest multiple of k that is greater or equal to x.
// k must be > 1
func (x Number) alignUp(k Number) (y Number) {
	x += (k - 1)
	y = x.align(k)
	return
}

// Min function
func (a Number) min(b Number) (y Number) {
	if a < b {
		y = a
	} else {
		y = b
	}
	return
}

// Max function
func (a Number) max(b Number) (y Number) {
	if a < b {
		y = b
	} else {
		y = a
	}
	return
}

func (x Number) isDerivedBy(k Number) bool {
	return x%k == 0
}

// Returns the 10 logarithm of x
func (x Number) log10() (y Number) {
	z := Number(10)
	for z <= x {
		z *= 10
		y += 1
	}
	return
}

// Used to sort result blocks using the first element
type ByFirst [][]Number

func (p ByFirst) Len() int           { return len(p) }
func (p ByFirst) Less(i, j int) bool { return p[i][0] < p[j][0] }
func (p ByFirst) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
