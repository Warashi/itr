package itr_test

import (
	"fmt"

	"github.com/Warashi/itr"
)

func ExampleSkipWhile() {
	it := func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
			fmt.Println("yield", i)
		}
	}

	skip := itr.SkipWhile(func(i int) bool { return i < 3 }, it)
	for i := range skip {
		fmt.Println("consume", i)
	}

	// Output:
	// yield 0
	// yield 1
	// yield 2
	// consume 3
	// yield 3
	// consume 4
	// yield 4
	// consume 5
	// yield 5
	// consume 6
	// yield 6
	// consume 7
	// yield 7
	// consume 8
	// yield 8
	// consume 9
	// yield 9
}

func ExampleSkipWhile_2() {
	it := func(yield func(int) bool) {
		for i := range 10 {
			if !yield(i) {
				return
			}
			fmt.Println("yield", i)
		}
	}

	skip := itr.SkipWhile(func(i int) bool { return i < 3 }, it)
	for i := range skip {
		fmt.Println("consume", i)
		break
	}

	// Output:
	// yield 0
	// yield 1
	// yield 2
	// consume 3
}
