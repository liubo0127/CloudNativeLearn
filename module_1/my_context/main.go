package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	baseCtx := context.Background()
	ctx := context.WithValue(baseCtx, "a", "b")
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(c context.Context) {
		fmt.Println(c.Value("a"))
		wg.Done()
	}(ctx)

	timeoutCtx, cancel := context.WithTimeout(baseCtx, time.Second)
	defer cancel()

	go func(c context.Context) {
		ticker := time.NewTicker(time.Second)
		for _ = range ticker.C {
			select {
			case <-c.Done():
				fmt.Println("child process interrupt")
				wg.Done()
				return
			default:
				fmt.Println("enter default")
			}
		}
	}(timeoutCtx)

	timer := time.NewTimer(time.Second)
	select {
	case <-timeoutCtx.Done():
		fmt.Println("main process exit!")
	case <-timer.C:
		fmt.Println("wait")
	}
	wg.Wait()
}
