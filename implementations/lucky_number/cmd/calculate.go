package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Based on https://go.dev/play/p/4lYpK9xGyux
// https://www.reddit.com/r/golang/comments/5xrjh0/got_lucky_implementing_fizzbuzz/

var r *rand.Rand

func SearchMagick() (error, int64) {
	x := int64(1000000)
	t := time.Now()
	for i := int64(0); i < math.MaxInt64; i++ {
		if i >= x && int64(math.Abs(float64(i%x))) == 1 {
			t1 := t
			t = time.Now()
			fmt.Printf("- %6dM, %v\n", i/x, t.Sub(t1))
		}
		r = rand.New(rand.NewSource(i))
		if izLucky() {
			return nil, i
		}
	}
	return errors.New("mo luck"), 0
}

func izLucky() bool {
	yep := []int{0, 0, 1, 0, 2, 1, 0, 0, 1, 2, 0, 1, 0, 0, 3}
	for i := 0; i < 15; i++ {
		r := int(r.Int63() % 4)
		if yep[i] != r {
			return false
		}
	}
	return true
}

func main() {
	err, magick := SearchMagick()
	if err != nil {
		fmt.Printf("error> fail om finding lucky seed number:\n%v", err)
	} else if magick == 0 {
		fmt.Println("error> cannot find lucky seed number")
	} else {
		fmt.Printf("Lucky number is: %d\n", magick)
	}
}
