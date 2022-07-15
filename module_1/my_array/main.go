package main

import "fmt"

func main() {
	var array1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
	fmt.Println(array1 == array2)
}
