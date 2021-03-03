package main

import "fmt"

func main() {
	favSport := "football"

	switch favSport {
	case "football":
		fmt.Println("⚽")
	case "basketball":
		fmt.Println("🏀")
	case "volleyball":
		fmt.Println("🏐")
	case "baseball":
		fmt.Println("⚾")
	}
}
