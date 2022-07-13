package main

import "fmt"

func forStandard(val int) int {
	// 标准 for 循环
	sum := 0
	for i := 1; i <= val; i++ {
		sum += i
	}
	return sum
}

func forOnlyCondition(val int) int {
	// 只带条件的 for 循环
	i := 1
	sum := 0
	for ;i <= val; {
		sum += i
		i++
	}
	return sum
}

func forInfinite(val int) int {
	// 无限循环，直到匹配条件退出
	i := 0

	for {
		if i > val {
			break
		}
		i++
	}
	return i
}

func forString(s string) {
	// 遍历字符串，遍历 map、数组、切片方法差不多
	for index, char := range s {
		fmt.Printf("index: %d, value: %v\n", index, string(char))
	}
}

func main() {
	sCnt := forStandard(10)
	oCnt := forOnlyCondition(10)
	iValue := forInfinite(10)
	fmt.Println(sCnt, oCnt, iValue)
	forString("Hello World!")
}
