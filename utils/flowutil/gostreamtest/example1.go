package main

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/17 17:14
 * @Desc:
 */

import (
	"fmt"
	"github.com/mariomac/gostream/stream"
)

/*
Example 1: basic creation, transformation and iteration

    Creates a literal stream containing all the integers from 1 to 11.
    From the Stream, selects all the integers that are prime
    For each filtered int, prints a message.

*/
func example1() {
	stream.Of(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11).
		Filter(isPrime).
		ForEach(func(n int) {
			fmt.Printf("%d is a prime number\n", n)
		})
}

func isPrime(n int) bool {
	for i := 2; i < n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
