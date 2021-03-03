package main

import "fmt"

func main() {
	myMap := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Printf("%T\n", myMap)
	fmt.Println(myMap)

	for i, v1 := range myMap {
		fmt.Println("Map:", i, v1)
		for j, v2 := range v1 {
			fmt.Println("\tSlice:", j, v2)
		}
	}
}
