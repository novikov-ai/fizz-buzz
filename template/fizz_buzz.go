package template

import (
	"fizz-buzz/pkg/contract"
)

type FizzBuzzer struct {
	contract.FizzBuzzer
	n int
}

func New(n int) contract.FizzBuzzer {
	return &FizzBuzzer{n: n}
}

func (fb *FizzBuzzer) FizzBuzz() []string {
	result := make([]string, 0, fb.n)

	// TODO: implement your solution here

	return result
}
