package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(context.Background(), "language", "Go")

	getValueFromContext := func(ctx context.Context, k string) {
		if v := ctx.Value(k); v != nil {
			fmt.Printf("Key - \"%v\" has the value of \"%v\"\n", k, v)
			return
		}
		fmt.Printf("Key - \"%v\" not found\n", k)
	}

	getValueFromContext(ctx, "language")
	getValueFromContext(ctx, "name")

	ctx = context.WithValue(ctx, "version", "1.16")
	fmt.Println("----------")
	getValueFromContext(ctx, "language")
	getValueFromContext(ctx, "name")
	getValueFromContext(ctx, "version")
}
