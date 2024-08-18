package itr

import (
	"iter"
)

// Skip skips the first n elements.
func Skip[T any](n int, it iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for v := range it {
			if i < n {
				i++
				continue
			}
			if !yield(v) {
				return
			}
		}
	}
}
