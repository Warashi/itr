package itr_test

import (
	"fmt"

	"github.com/Warashi/itr"
)

func ExampleTakeWhile() {
	it := func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
			fmt.Println("yield", i)
		}
	}

	take := itr.TakeWhile(func(i int) bool { return i < 3 }, it)
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

func ExampleTakeWhile_2() {
	it := func(yield func(int) bool) {
		for i := range 3 {
			if !yield(i) {
				return
			}
			fmt.Println("yield", i)
		}
	}

	take := itr.TakeWhile(func(i int) bool { return i < 10 }, it)
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
