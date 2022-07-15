package main

import "fmt"

type personInfo struct {
	age     string
	address string
}

func (p *personInfo) getAddress() (address string) {
	address = p.address
	return
}

func (p *personInfo) getAge() (age string) {
	age = p.age
	return
}

func main() {
	var myMap1 = map[string]string{
		"name":    "aaa",
		"address": "北京",
	}
	fmt.Println(myMap1["name"], myMap1["address"])

	myMap2 := make(map[string]string, 3) // 会根据数据扩大长度，内存地址不变
	myMap2["name"] = "aaa"
	myMap2["address"] = "北京"
	myMap2["comment"] = "test"
	myMap3 := myMap2
	myMap2["extra"] = "a"
	fmt.Println(myMap2, len(myMap2), myMap3) // myMap2 新增第四个字段之后，myMap3 的长度也变成了 4，说明 myMap2 和 myMap3 的内存地址在扩容之后也是一致的

	myMap4 := make(map[string]personInfo)
	myMap4["zhangsan"] = personInfo{
		age:     "18",
		address: "北京",
	}
	myMap4["lisi"] = personInfo{
		age:     "20",
		address: "上海",
	}
	fmt.Println(myMap4)

	person1, exists := myMap4["zhangsan"]
	if !exists {
		fmt.Println("值不存在")
	} else {
		person1.getAge()
		fmt.Println(person1.getAge(), person1.getAddress())
	}

	for name, info := range myMap4 {
		fmt.Printf("name: %s, age: %s, address: %s\n", name, info.getAge(), info.getAddress())
	}
}
