# Lucky Number implementation of FizzBuzz

Based on searching Google for tricks helping find shortest 
possible Go code for solve FizzBuzz problem. 

## Lucky number

Method based on finding lucky number which gives necessary index if get module of 4.
This index (0..3) can be used for get correct combination.

See / run [calculation](./cmd/calculate.go) utility.
On local computer it takes less than 6 hours for find this number for Go rand.

At version Go 1.23.2 lucky number is: `176064004`

## Links
- Go Playground: https://go.dev/play/p/4lYpK9xGyux
- Reddit: https://www.reddit.com/r/golang/comments/5xrjh0/got_lucky_implementing_fizzbuzz/
