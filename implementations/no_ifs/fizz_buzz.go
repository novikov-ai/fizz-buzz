package no_ifs

import (
	"fmt"

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

	f := []string{"Fizz", "", ""}
	b := []string{"Buzz", "", "", "", ""}

	for i := 1; i < fb.n; i++ {
		t := f[i%3] + b[i%5]
		opts := []interface{}{i, t, t, t, t, t, t, t, t}
		printed := fmt.Sprintf("%v", opts[len(t)])

		result = append(result, printed)
	}

	return result
}
