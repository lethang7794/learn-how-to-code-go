package main

import "fmt"

type person struct {
	first           string
	last            string
	iceCreamFlavors []string
}

func main() {
	p1 := person{
		first: "Thang",
		last:  "Le",
		iceCreamFlavors: []string{
			"chocolate",
			"lemon",
			"almond",
		},
	}

	p2 := person{
		first: "Lieu",
		last:  "Truong",
		iceCreamFlavors: []string{
			"lemon",
			"lemon",
			"lemon",
		},
	}

	p := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(p)
	fmt.Printf("%T\n", p)

	for i, v := range p {
		fmt.Println(i, v)
		fmt.Println("First:", v.first)
		fmt.Println("Last:", v.last)
		fmt.Println("Ice Cream Flavors:")
		for j, v2 := range v.iceCreamFlavors {
			fmt.Printf("\t%v\t%v\n", j, v2)
		}
	}
}
