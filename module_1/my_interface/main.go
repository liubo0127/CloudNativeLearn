package main

import "fmt"

type myInterface interface {
	getName() string
}

type Human struct {
	firstName string
	lastName  string
}

type Plane struct {
	vendor string
	model  string
}

func (h *Human) getName() string {
	return fmt.Sprintf("%s,%s", h.firstName, h.lastName)
}

func (p *Plane) getName() string {
	return fmt.Sprintf("vendor: %s, model: %s", p.vendor, p.model)
}

func main() {
	interface_slice := []myInterface{}
	h := new(Human)
	h.firstName = "first"
	h.lastName = "last"
	fmt.Println(h.getName())
	p := &Plane{}
	p.model = "testModel"
	p.vendor = "testVendor"
	fmt.Println(p.getName())
	interface_slice = append(interface_slice, h)
	interface_slice = append(interface_slice, p)
	for _, value := range interface_slice {
		fmt.Println(value.getName())
	}
}
