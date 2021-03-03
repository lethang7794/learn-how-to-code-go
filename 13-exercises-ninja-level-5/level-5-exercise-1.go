package main

import "fmt"

type person struct {
	first           string
	last            string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		first:           "Thang",
		last:            "Le",
		iceCreamFlavors: []string{"chocolate", "lemon", "almond"},
	}

	p2 := person{
		first:           "Lieu",
		last:            "Truong",
		iceCreamFlavors: []string{"lemon", "lemon", "lemon"},
	}

	fmt.Println(p1)
	for _, v := range p1.iceCreamFlavors {
		fmt.Println(v)
	}

	fmt.Println(p2)
	for _, v := range p2.iceCreamFlavors {
		fmt.Println(v)
	}
}
