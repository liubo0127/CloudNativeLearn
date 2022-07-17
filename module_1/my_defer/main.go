package main

import (
	"fmt"
	"sync"
)

func myDefer(a, b int) int {
	defer fmt.Println("Im first defer")
	defer fmt.Println("Im second defer")

	c := a + b
	fmt.Printf("c value: %d\n", c)
	return c
}

func loopFunction() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			fmt.Println("loopfuntion: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func main() {
	fmt.Println(myDefer(1, 2))

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	loopFunction()
}
