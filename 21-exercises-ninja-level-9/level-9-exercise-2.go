package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("Hello from", p.name)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		name: "Thang",
		age:  27,
	}

	saySomething(&p1)
	// saySomething(p1) // This line won't work.
}
