package main

import (
	"context"
	"fmt"
)

func main() { //O context pode ser usado para tranferencia de informações
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	token := ctx.Value("token")
	fmt.Println(token)
}
