package main

import "fmt"

func main() {
	name := "testing"
	fmt.Printf("%s\n", name)
	fmt.Printf("%d\n", name)		//类型不匹配，go vet 会提示异常
}
