//go:build ignore

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// HTTPレスポンスとして、 "Hi there, I love %s!"をクライアントへ送信する
	// r.URL.Path[1:]はリクエストの先頭の"/"を除外した、URLパスの残りの部分
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	// http.ListenAndServeでHTTPサーバーをポート8080で起動する
	// サーバーでエラーが発生した場合は、log.Fatal()がエラーログを出力し、プログラムを終了させる
	log.Fatal(http.ListenAndServe(":8080", nil))
}
