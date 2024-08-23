package main

import (
	"context"
	"fmt"
)

type TraceID string

const ZeroTraceID = ""

// キーは空構造体を用いて独自型を定義するのが一般的
// プリミティブ型は他のパッケージと衝突する可能性があるので避ける
type traceIDKey struct{}

// context.Context型に値をセットする
func SetTraceID(ctx context.Context, tid TraceID) context.Context {
	return context.WithValue(ctx, traceIDKey{}, tid)
}

// context.Context型に値にセットされた値を取得する
func GetTraceID(ctx context.Context) TraceID {
	if v, ok := ctx.Value(traceIDKey{}).(TraceID); ok {
		return v
	}
	return ZeroTraceID
}

func main() {
	ctx := context.Background()
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
	ctx = SetTraceID(ctx, "test-id")
	fmt.Printf("trace id = %q\n", GetTraceID(ctx))
}
