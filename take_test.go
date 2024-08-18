package itr_test

import (
	"fmt"

	"github.com/Warashi/itr"
)

func ExampleTake() {
	it := func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
			fmt.Println("yield", i)
		}
	}

	take := itr.Take(3, it)
	for i := range take {
		fmt.Println("consume", i)
	}

	// Output:
	// consume 0
	// yield 0
	// consume 1
	// yield 1
	// consume 2
	// yield 2
}

func ExampleTake_2() {
	it := func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
			fmt.Println("yield", i)
		}
	}

	take := itr.Take(3, it)
	for i := range take {
		fmt.Println("consume", i)
		break
	}

	// Output:
	// consume 0
}
