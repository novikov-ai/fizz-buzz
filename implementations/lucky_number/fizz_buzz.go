package lucky_number

import (
	"fizz-buzz/pkg/contract"
	"math/rand"
	"strconv"
)

const (
	LUCKY = 176064004
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

	predefined := []string{"", "Fizz", "Buzz", "FizzBuzz"}
	var r *rand.Rand
	for i := 1; i < fb.n; i++ {
		if i%15 == 1 {
			r = rand.New(rand.NewSource(LUCKY))
		}
		index := r.Int63() % 4
		if index == 0 {
			result = append(result, strconv.Itoa(i))
		} else {
			result = append(result, predefined[index])
		}
	}

	return result
}
