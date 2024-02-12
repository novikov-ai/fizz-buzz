package main

import (
	no_ifs "fizz-buzz/implementations/no-ifs"
)

func main() {
	// EXAMPLE
	for _, s := range no_ifs.New(100).FizzBuzz() {
		println(s)
	}
}
