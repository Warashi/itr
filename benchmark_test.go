package itr_test

import (
	"iter"
	"testing"

	"github.com/Warashi/itr"
)

func BenchmarkPrimes(b *testing.B) {
	var (
		primes  iter.Seq[int]
		isPrime func(int) bool
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

	primes = func(yield func(int) bool) {
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
	}

	for _ = range itr.Take(100000, primes) {

	}
}

func BenchmarkCachePrimes(b *testing.B) {
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

	primes = func(yield func(int) bool) {
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
	}

	primes, cleanup = itr.Cache(primes)
	defer cleanup()

	for _ = range itr.Take(100000, primes) {

	}
}
