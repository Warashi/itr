package itr_test

import (
	"fmt"
	"iter"

	"github.com/Warashi/itr"
)

func Example() {
	var (
		primes  iter.Seq[int]
		isPrime func(int) bool
		cleanup func()
	)

	isPrime = func(n int) bool {
		for p := range primes {
			if p*p > n {
				return true
			}
			if n%p == 0 {
				return false
			}
		}
		return true
	}

	primes, cleanup = itr.Cache(func(yield func(int) bool) {
		if !yield(2) {
			return
		}
		if !yield(3) {
			return
		}
		for i := 5; ; i += 2 {
			if isPrime(i) {
				if !yield(i) {
					return
				}
			}
		}
	})

	defer cleanup()

	for i := range itr.Take(10, primes) {
		fmt.Println(i)
	}

	// Output:
	// 2
	// 3
	// 5
	// 7
	// 11
	// 13
	// 17
	// 19
	// 23
	// 29
}
