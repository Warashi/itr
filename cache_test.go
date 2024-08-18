package itr_test

import (
	"fmt"
	"iter"
	"testing"

	"github.com/Warashi/itr"
)

func ExampleCache(t *testing.T) {
	it := func() iter.Seq[int] {
		return func(yield func(int) bool) {
			for i := range 3 {
				fmt.Println("yield", i)
				yield(i)
			}
		}
	}

	cached, stop := itr.Cache(it())
	t.Cleanup(stop)

	for i := range cached {
		fmt.Println("consume", i)
	}
	for i := range cached {
		fmt.Println("consume", i)
	}

	// Output:
	// yield 0
	// consume 0
	// yield 1
	// consume 1
	// yield 2
	// consume 2
	// consume 0
	// consume 1
	// consume 2
}
