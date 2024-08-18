package itr

import (
	"iter"
)

func Take[T any](n int, it iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		i := 0
		for v := range it {
			if i >= n {
				return
			}
			if !yield(v) {
				return
			}
			i++
		}
	}
}
