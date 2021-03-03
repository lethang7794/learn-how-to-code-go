package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "Thang",
		age:  27,
	}
	fmt.Println(p1)

	//! A reference to a non-interface method with a pointer receiver using an addressable value
	//! will automatically take the address of that value: t.Mp is equivalent to (&t).Mp.
	p1.changeMeMethod("Lieu", 27) // ~= (&p1).changeMeMethod("Lieu", 27)
	fmt.Println(p1)
	(&p1).changeMeMethod("Harry", 20)
	fmt.Println(p1)

	changeMe(&p1)
	fmt.Println(p1)
}

func (p *person) changeMeMethod(name string, age int) {
	(*p).name = name // Dereferencing the struct pointer.
	//! A reference to a non-interface method with a value receiver using a pointer
	//! will automatically dereference that pointer: pt.Mv is equivalent to (*pt).Mv.
	p.age = age // ~= (*p).age = age <- Struct pointer are automatically dereferenced.
}

func changeMe(p *person) {
	(*p).name = "James"
	(*p).age = 35
}
