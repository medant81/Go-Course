/*
Нужно создать контекст со значениями

some key1: some value1
some key2: some value2

Контекст следует передать в функцию, которая выведет значения some key1 и some key2 в stdout.
*/
package main

import (
	"context"
	"fmt"
)

type ctxKey string

func main() {

	ctx := context.Background()

	var key ctxKey = "some key1"
	ctx = context.WithValue(ctx, key, "some value1")

	key = "some key2"
	ctx = context.WithValue(ctx, key, "some value2")

	do(ctx)

}

func do(ctx context.Context) {

	var key ctxKey = "some key1"

	fmt.Printf("%s: %s\n", key, ctx.Value(key))
	key = "some key2"
	fmt.Printf("%s: %s\n", key, ctx.Value(key))

}
