package main

import (
	"fmt"
	"go.uber.org/ratelimit"
)

func RateLimit() {
	r := ratelimit.New(10)

	for i := 0; i < 100; i++ {
		a := r.Take()
		fmt.Println(i, a)
	}
}
