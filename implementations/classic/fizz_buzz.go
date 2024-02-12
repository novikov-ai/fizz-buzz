package classic

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
	result := make([]string, 0, fb.n)

	for i := 1; i < fb.n; i++ {
		printed := ""
		if i%3 == 0 {
			printed += "Fizz"
		}
		if i%5 == 0 {
			printed += "Buzz"
		}

		if printed == "" {
			printed = strconv.Itoa(i)
		}

		result = append(result, printed)
	}

	return result
}
