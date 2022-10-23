package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func main() {

	// ハンドラファンクション -> デフォルトmultiplexerに登録。
	// 第一引数 -> パターン
	// 第二引数 -> 「レスポンスライター」「リクエスト」を引数として受け取る関数
	http.HandleFunc("/greet", greet)

	// 第一引数 -> リッスンするアドレス
	// 第二引数 -> 使用するmultiplexer | デフォルトで用意されているものを使用するため、nilを設定
	http.ListenAndServe("localhost:8000", nil)
}
