package main

import "fmt"

func main() {
	fmt.Println("panic test")
	defer func() {
		fmt.Println("defer function is called")
		if err := recover(); err != nil {
			fmt.Println(err) // err = "a panic is triggered"
		}
	}()
	panic("a panic is triggered")
}
