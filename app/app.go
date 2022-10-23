package app

import "net/http"

func Start() {
	// 自身でmultiplexerを登録
	mux := http.NewServeMux()

	// ハンドラファンクション -> デフォルトmultiplexerに登録。 | 自身で作成したmultiplexerを使用
	// 第一引数 -> パターン
	// 第二引数 -> 「レスポンスライター」「リクエスト」を引数として受け取る関数
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", getAllCustomers)
	mux.HandleFunc("/customers_xml", getAllCustomers_xml)
	mux.HandleFunc("/customers_flex", getAllCustomers_flex)

	// 第一引数 -> リッスンするアドレス
	// 第二引数 -> 使用するmultiplexer | デフォルトで用意されているものを使用するため、nilを設定
	http.ListenAndServe("localhost:8000", mux)
}
