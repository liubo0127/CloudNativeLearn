package main

import (
	"flag"
	"fmt"
	"os"
)

// 最简单的传参方法，第一个参数为代码文件本身，从第二个开始是真正传的参数
func getArgs() []string {
	return os.Args
}

func getFlags() {
	name := flag.String("name", "world", "test")
	flag.Parse() // 解析参数，不解析只有默认值
	fmt.Printf("name: %T, %s\n", *name, *name)

}

func returnMultiValues(a, b int) (x, y int) {
	x, y = a, b
	return
}

func main() {
	//parameters := getArgs()
	//fmt.Printf("args: %v\n", parameters)
	getFlags()
	fmt.Println(os.Args)
	fmt.Println(returnMultiValues(1, 2))
}
