package main

import (
	"fmt"
	"sync"
	"time"
)

func waitBySleep() {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
	time.Sleep(time.Second)
}

func waitByChannel() {
	ch := make(chan bool, 10)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			ch <- true
		}(i)
	}

	cnt := 0
	for i := 0; i < 100; i++ {
		<-ch
		cnt++
	}
	fmt.Println(cnt)
}

func waitbyWg() {
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

func main() {
	//waitBySleep()
	//waitByChannel()
	waitbyWg()
}
