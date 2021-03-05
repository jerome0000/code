package gocode

import (
	"context"
	"fmt"
	"time"
)

func ctx() {
	// var wg sync.WaitGroup
	// wg.Add(2)
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("wg 1 ok")
	// 	wg.Done()
	// }()
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("wg 2 ok")
	// 	wg.Done()
	// }()
	// wg.Wait()

	// stop := make(chan bool)
	// go func() {
	// 	for {
	// 		select {
	// 		case <-stop:
	// 			fmt.Println("sia end...")
	// 			return
	// 		default:
	// 			fmt.Println("sia doing")
	// 			time.Sleep(1 * time.Second)
	// 		}
	// 	}
	// }()
	// time.Sleep(10 * time.Second)
	// fmt.Println("sia can end")
	// stop <- true
	// time.Sleep(5 * time.Second)
	//
	// ctx, cancel := context.WithCancel(context.Background())
	// go func(ctx context.Context) {
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println("sia end...")
	// 			return
	// 		default:
	// 			fmt.Println("sia doing")
	// 			time.Sleep(1 * time.Second)
	//
	// 		}
	// 	}
	// }(ctx)
	// time.Sleep(10 * time.Second)
	// fmt.Println("sia can end")
	// cancel()

	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "sia 1")
	go watch(ctx, "sia 2")
	go watch(ctx, "sia 3")
	time.Sleep(1 * time.Second)
	fmt.Println("sia do end")
	cancel()
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, ", end")
			return
		default:
			fmt.Println(name, ", doing")
			time.Sleep(2 * time.Second)
		}
	}
}
