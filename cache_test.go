package itr_test

import (
	"fmt"
	"iter"

	"github.com/Warashi/itr"
)

func ExampleCache() {
	it := func() iter.Seq[int] {
		return func(yield func(int) bool) {
			for i := range 3 {
				fmt.Println("yield", i)
				if !yield(i) {
					return
				}
			}
		}
	}

	cached, stop := itr.Cache(it())
	defer stop()

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
