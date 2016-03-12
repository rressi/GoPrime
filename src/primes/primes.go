package primes

import (
	"sync"
)

type Number uint64

// Returns the squared root of a number strictly positive
func (x Number) sqrt() (y Number) {

	for y = 1; y*y <= x; y++ {
	}
	y--

	// fmt.Println("sqrt(", x, ")")
	// fmt.Println("    y", y)
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

// Returns the 10 logarithm of x
func (x Number) log10() (y Number) {
	z := Number(10)
	for z <= x {
		z *= 10
		y += 1

	}
	return
}


// Finds prime numbers below N.
func FindPrimes(N Number) (primes []Number) {

	capacity := Number(20)
	if N > 100 {
		capacity = N / N.log10()
	}
	// fmt.Println(N, "->", capacity)

	primes = make([]Number, 0, capacity)
	if N <= 2 {
		return
	}

	nonPrimes := make([]bool, N)
	findNonPrimes(nonPrimes)

	var j Number
	primes = append(primes, 2)
	j++

	X := Number(len(nonPrimes))
	for x := Number(3); x + 1 <= X; x += 2 {
		if !nonPrimes[x] {
			primes = append(primes, x)
			j++
		}
	}

	// fmt.Println("primes", primes)
	return
}

func findNonPrimes(nonPrimes []bool) {

	N := Number(len(nonPrimes))
	// fmt.Println("findNonPrimes(", N, nonPrimes, ")")
	// defer func() {
	// 	fmt.Println("    nonPrimes", nonPrimes)
	// }()

	// First non primes are hard-coded:
	if N <= 16 {
		copy(nonPrimes, []bool{
			true, true, false, false,    // 0 - 3
			true, false, true, false,    // 4 - 7
			true, true, true, false,     // 8 - 11
			true, false, true, true,})   // 12 - 15
		return
	}

	// Recursively finds non-primes below or equal to sqrt(N)
	n := N.sqrt() + 1
	findNonPrimes(nonPrimes[:n])

	if n < 1000 {
		// Fills serially the buffer of non primes:
		for x1 := n; x1 < N; x1 += n {
			x2 := x1 + n
			if x2 > N {
				x2 = N
			}
			fillNonPrimes(x1, nonPrimes[:n], nonPrimes[x1:x2])
		}
	} else {
		// Fills on parallel the buffer of non primes:
		var wg sync.WaitGroup
		for x1 := n; x1 < N; x1 += n {
			x2 := x1 + n
			if x2 > N {
				x2 = N
			}
			go func(x1, x2 Number) {
				fillNonPrimes(x1, nonPrimes[:n], nonPrimes[x1:x2])
				wg.Done()
			}(x1, x2)
			wg.Add(1)
		}
		wg.Wait()
	}

	return
}

func fillNonPrimes(yStart Number, in, out []bool) {

	// fmt.Println("fillNonPrimes(", yStart, in, out, ")")
	// defer func() {
	// 	fmt.Println("    out", out)
	// }()

	X := Number(len(in))
	Y := Number(len(out))

	// NOTE: we start from 3 because we ignore even numbers.
	for x := Number(3); x < X; x++ {
		if !in[x] {
			y := yStart.alignUp(x) - yStart
			for ; y < Y; y += x {
				out[y] = true
			}
		}
	}

	return
}
