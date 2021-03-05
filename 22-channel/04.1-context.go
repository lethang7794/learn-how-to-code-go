package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Printf("Context value: %v\t\n", ctx)
	fmt.Printf("Context type : %T\t\n", ctx)
	fmt.Printf("Context err  : %v\t\n", ctx.Err())

	ctx2, cancel := context.WithCancel(ctx)
	fmt.Println("-------------")
	fmt.Printf("Context value: %v\t\n", ctx2)
	fmt.Printf("Context type : %T\t\n", ctx2)
	fmt.Printf("Context err  : %v\t\n", ctx2.Err())

	cancel()
	fmt.Println("-------------")
	fmt.Printf("Context value: %v\t\n", ctx)
	fmt.Printf("Context type : %T\t\n", ctx)
	fmt.Printf("Context err  : %v\t\n", ctx.Err())
	fmt.Println("-------------")
	fmt.Printf("Context value: %v\t\n", ctx2)
	fmt.Printf("Context type : %T\t\n", ctx2)
	fmt.Printf("Context err  : %v\t\n", ctx2.Err())
}
