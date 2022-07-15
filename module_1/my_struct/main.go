package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	Name string `string:"name"`
	Age  int    `int:"age"`
}

func printMyStruct(m *myStruct) {
	fmt.Printf("type: %T, value: %v\n", m.Name, m.Name)
}

func main() {
	m := myStruct{Name: "zhangsan"}
	fmt.Println(m, m.Name)
	printMyStruct(&m)
	reflectM := reflect.TypeOf(m)
	col1 := reflectM.Field(0)              // 取 struct 的第 n 个字段，0 表示第一个字段
	col2, _ := reflectM.FieldByName("Age") // 取 struct 中指定名字的字段 （不是标签名字）
	fmt.Println(col2.Name)                 // 获取struct 字段的名称
	fmt.Println(col1.Tag.Get("string"), col1.Name)
	fmt.Println(col2.Tag.Get("int"))
}
