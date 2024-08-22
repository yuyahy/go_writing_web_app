package main

import (
	"context"
	"fmt"
)

func child(ctx context.Context) {

	if err := ctx.Err(); err != nil {
		return
	}
	fmt.Println("Not cancel")
}

func main() {
	// WithCancelで生成したcontext.Contextを関数間で共有することで、
	// contextがキャンセルされたら関係する処理を全てキャンセルすることができる
	ctx, cancel := context.WithCancel(context.Background())
	child(ctx)
	cancel()
	child(ctx)
}
