package itr_test

import (
	"fmt"
	"testing"

	"github.com/Warashi/itr"
)

func ExampleTake(t *testing.T) {
	it := func(yield func(int) bool) {
		for i := range 10 {
			fmt.Println("yield", i)
			yield(i)
		}
	}

	take := itr.Take(3, it)
	for i := range take {
		fmt.Println("consume", i)
	}

	// Output:
	// yield 0
	// consume 0
	// yield 1
	// consume 1
	// yield 2
	// consume 2
}
