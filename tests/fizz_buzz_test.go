package tests

import (
	"fizz-buzz/implementations/classic"
	noIfs "fizz-buzz/implementations/no-ifs"
	"fizz-buzz/pkg/contract"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const maxNumber = 100

const output = `1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
Fizz
19
Buzz
Fizz
22
23
Fizz
Buzz
26
Fizz
28
29
FizzBuzz
31
32
Fizz
34
Buzz
Fizz
37
38
Fizz
Buzz
41
Fizz
43
44
FizzBuzz
46
47
Fizz
49
Buzz
Fizz
52
53
Fizz
Buzz
56
Fizz
58
59
FizzBuzz
61
62
Fizz
64
Buzz
Fizz
67
68
Fizz
Buzz
71
Fizz
73
74
FizzBuzz
76
77
Fizz
79
Buzz
Fizz
82
83
Fizz
Buzz
86
Fizz
88
89
FizzBuzz
91
92
Fizz
94
Buzz
Fizz
97
98
Fizz
Buzz
`

func Test_FizzBuzz(t *testing.T) {

	tests := []struct {
		name string
		do   contract.FizzBuzzer
	}{
		{
			name: "classic",
			do:   classic.New(maxNumber),
		},
		{
			name: "without if",
			do:   noIfs.New(maxNumber),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, fb := range tt.do.FizzBuzz() {
				output := strings.Split(output, "\n")
				assert.Equal(t, output[i], fb)
			}
		})
	}
}
