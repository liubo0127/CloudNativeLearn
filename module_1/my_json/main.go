package main

import (
	"encoding/json"
	"fmt"
)

type Human struct {
	Name    string `string:name`
	Age     int    `int:age`
	Address string `string:address`
}

func unmarshal2Struct(s string) Human {
	h := Human{}
	err := json.Unmarshal([]byte(s), &h)
	if err != nil {
		fmt.Println(err)
	}
	return h
}

func marshal2String(h *Human) string {
	marshal, err := json.Marshal(h)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(marshal)
}

func main() {
	s := `{"name":"a","age":18,"address":"bj"}`
	h := unmarshal2Struct(s)
	fmt.Println(h)
	ss := marshal2String(&h)
	fmt.Println(ss)
}
