package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	aTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}

	aSedan := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: true,
	}

	fmt.Println(aTruck)
	fmt.Println(aTruck.fourWheel)
	fmt.Println(aTruck.doors)
	fmt.Println(aTruck.color)

	fmt.Println(aSedan)
	fmt.Println(aSedan.luxury)
	fmt.Println(aTruck.doors)
	fmt.Println(aTruck.color)
}
