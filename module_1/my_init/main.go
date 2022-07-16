package main

import (
	"fmt"

	_ "github.com/liubo0127/CloudNativeLearn/module_1/my_init/init_a"
	_ "github.com/liubo0127/CloudNativeLearn/module_1/my_init/init_b"
)

func init() {
	fmt.Println("Main init")
}

func main() {
	fmt.Println("Main function")
}
