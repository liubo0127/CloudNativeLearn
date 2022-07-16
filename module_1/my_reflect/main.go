package main

import (
	"fmt"
	"reflect"
)

type Human struct {
	name    string
	age     int
	address string
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetAge() int {
	return h.age
}
func (h *Human) GetAddress() string {
	return h.address
}

//func (h *Human) setAge(age int) {
//	h.age = age
//}

func main() {
	m := map[string]string{
		"a": "b",
	}
	mType := reflect.TypeOf(m)
	mValue := reflect.ValueOf(m)
	fmt.Println(mType, mValue)

	s := &Human{
		name:    "zhangsan",
		age:     18,
		address: "北京",
	}
	fmt.Println(s.GetAge())
	sType := reflect.TypeOf(s)
	sValue := reflect.ValueOf(s)
	fmt.Println(sType, sValue)

	fmt.Println(sValue.Interface().(*Human))      //将变量转换成 struct
	fmt.Println(sValue.NumMethod())               // 变量的方法数量
	fmt.Println(sValue.Elem().FieldByName("age")) // Elem() 获取指针对应的值
	for i := 0; i < sValue.NumMethod(); i++ {
		t := sValue.Method(i)
		fmt.Println(t.Call(nil))
	}
}
