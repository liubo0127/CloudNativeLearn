package main

import "fmt"

func ifStandard(val int) int {
	/* 标准示例 */
	if val > 10 {
		return 10
	} else if val > 5 {
		return 5
	} else {
		return 0
	}
}

func ifAssignValue(val int) int {
	if i := val; i > 10 {
		return 10
	} else if i > 5 {
		return 5
	} else {
		return 0
	}
}

func main() {
	s := ifStandard(11)
	a := ifAssignValue(11)
	fmt.Printf("s: %d, a: %d\n", s, a)
}
