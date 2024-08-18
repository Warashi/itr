package itr

import (
	"iter"
)

func TakeWhile[T any](pred func(T) bool, it iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range it {
			if !pred(v) {
				return
			}
			if !yield(v) {
				return
			}
		}
	}
}
