package gocode

import (
	"fmt"
	"time"
)

var POOL = 100

func groutine1(p chan int) {
	for i := 1; i <= POOL; i++ {
		p <- i
		if i%2 == 1 {
			fmt.Println("groutine-1:", i)
		}
	}
}

func groutine2(p chan int) {
	for i := 1; i <= POOL; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("groutine-2:", i)
		}
	}
}

func printNum() {
	ch := make(chan int)
	go groutine1(ch)
	go groutine2(ch)

	time.Sleep(time.Second * 2)
}
