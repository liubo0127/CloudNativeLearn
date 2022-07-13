package main

import "fmt"

func switchStandard(val int) {
	switch val {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
		fallthrough
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println(999999999)
	}
}

func main() {
	val := 2
	switchStandard(val)
}