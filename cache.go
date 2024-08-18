package itr

import "iter"

// Cache caches the elements of the iterator.
func Cache[T any](it iter.Seq[T]) (_ iter.Seq[T], cleanup func()) {
	var c []T
	next, stop := iter.Pull(it)

	return func(yield func(T) bool) {
		for _, v := range c {
			if !yield(v) {
				return
			}
		}
		for {
			v, ok := next()
			if !ok {
				return
			}
			c = append(c, v)
			if !yield(v) {
				return
			}
		}
	}, stop
}
