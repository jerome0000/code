package gocode005

import (
	"context"
	"fmt"
	"time"
)

func hardWork(job interface{}) error {
	time.Sleep(time.Second * 10)
	fmt.Println("this is test")
	return nil
}

func requestWork(ctx context.Context, job interface{}) error {

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- hardWork(job)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// func main() {
// 	ctx := context.Background()
// 	err := requestWork(ctx, nil)
// 	fmt.Println(err.Error())
// }
