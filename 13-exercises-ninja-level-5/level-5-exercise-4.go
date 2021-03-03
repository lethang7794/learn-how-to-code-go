package main

import "fmt"

func main() {
	fish := struct {
		legs    int
		canFly  bool
		canSwim bool
	}{
		legs:    0,
		canSwim: true,
	}

	fmt.Println(fish)
}
