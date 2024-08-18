package itr

import (
	"iter"
)

// SkipWhile skips elements while the predicate is true.
func SkipWhile[T any](pred func(T) bool, it iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		skip := true
		for v := range it {
			if skip && pred(v) {
				continue
			}
			skip = false
			if !yield(v) {
				return
			}
		}
	}
}
