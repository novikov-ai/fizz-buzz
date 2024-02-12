package tests

import (
	"fizz-buzz/implementations/classic"
	"fizz-buzz/implementations/no_ifs"
	"fizz-buzz/pkg/contract"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
			do:   no_ifs.New(maxNumber),
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
