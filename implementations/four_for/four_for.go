package four_for

import (
	"fizz-buzz/pkg/contract"
	"strconv"
)

type FizzBuzzer struct {
	contract.FizzBuzzer
	n int
}

func New(n int) contract.FizzBuzzer {
	return &FizzBuzzer{n: n}
}

func (fb *FizzBuzzer) FizzBuzz() []string {
	result := make([]string, fb.n)

	for i := 0; i < fb.n; i++ {
		result[i] = strconv.Itoa(i + 1)
	}

	for i := 2; i < fb.n; i += 3 {
		result[i] = "Fizz"
	}

	for i := 4; i < fb.n; i += 5 {
		result[i] = "Buzz"
	}

	for i := 14; i < fb.n; i += 15 {
		result[i] = "FizzBuzz"
	}

	return result
}
