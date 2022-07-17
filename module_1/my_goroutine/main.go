package main

import (
	"fmt"
	"math/rand"
	"time"
)

func helloGoroutine() {
	go fmt.Println(1)
	go fmt.Println(2)
	go fmt.Println(3)
	time.Sleep(time.Second)
}

// 无缓冲通道
func myNoBufferChannel() {
	ch := make(chan int)
	go func() {
		fmt.Println("[No buffer channel] send data to channel")
		ch <- 1
	}()
	a := <-ch
	fmt.Println("[No buffer channel]", a)
}

// 缓冲通道
func myBufferChannel() {
	ch := make(chan int, 10)
	go func() {
		fmt.Println("[Buffer channel] Starting put data to channel")
		for i := 0; i < 10; i++ {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(10) // 0 <= n <= 10
			fmt.Println("[Buffer channel] putting: ", n)
			ch <- n
		}
		close(ch)
	}()

	fmt.Println("[Buffer channel] Starting get data from channel")
	for n := range ch {
		fmt.Println("[Buffer channel] getting: ", n)
	}
}

// 关闭通道
func channelClosed() {
	ch := make(chan int)

	go func() {
		fmt.Println("[channel closed] put data")
		ch <- 1
		close(ch)
	}()

	if v, notClosed := <-ch; notClosed {
		fmt.Println(v)
	}
}

// select test
func mySelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("[channel1] put data: ", 1)
			//time.Sleep(time.Second)
			ch1 <- 1
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("[channel2] put data: ", 2)
			//time.Sleep(time.Second)
			ch2 <- 2
		}
	}()
	timer := time.NewTimer(time.Second)
	for {
		select {
		case v := <-ch1:
			fmt.Println("[channel1] get data: ", v)
		case v := <-ch2:
			fmt.Println("[channel2] get data: ", v)
		//default:
		//	fmt.Println("[no data]")
		case <-timer.C:
			fmt.Println("[no data] timeout")
		}
		time.Sleep(time.Second)
	}
}

func main() {
	helloGoroutine()
	myNoBufferChannel()
	myBufferChannel()
	channelClosed()
	mySelect()
}
