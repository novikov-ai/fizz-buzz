package main

import (
	"fizz-buzz/implementations/no_ifs"
)

func main() {
	// EXAMPLE
	for _, s := range no_ifs.New(100).FizzBuzz() {
		println(s)
	}
}
