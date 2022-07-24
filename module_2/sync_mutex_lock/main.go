package main

import (
	"fmt"
	"sync"
	"time"
)

func rLock() {
	rlock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		rlock.RLock()
		defer rlock.RUnlock()
		fmt.Printf("[Rlock] %d\n", i)
	}
}

func wLock() {
	wlock := sync.RWMutex{}
	for i := 0; i < 3; i++ {
		wlock.Lock()
		defer wlock.Unlock()
		fmt.Printf("[Wlock] %d\n", i)
	}
}

func lock() {
	lock := sync.Mutex{}

	for i := 0; i < 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Printf("[Lock] %d\n", i)
	}
}

func main() {
	go rLock()
	go wLock()
	go lock()
	time.Sleep(time.Second * 5)
}
