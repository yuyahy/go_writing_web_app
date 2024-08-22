package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	go func() { fmt.Println("Other goroutine") }()
	fmt.Println("STOP")
	// context.Context.Done()で通知されるキャンセルまで待機する
	<-ctx.Done()
	fmt.Println("Then, the time has begun")
}
