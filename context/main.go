package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("We changed Key value to %s.\n", ctx.Value("Key"))
}

func main() {
	ctx := context.TODO()

	ctx = context.WithValue(ctx, "Key", "KeyValue")
	doSomething(ctx)
}
