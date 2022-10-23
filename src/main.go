package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON解析用に使用する構造体の定義
// JSON出力用のキーは「`json:★★★`」と指定することが可能。
// 指定しなければ、構造体のキーがそのまま使用される。
type Customer struct {
	Name string `json:fill_name`
	City string `json:city`
	Zipcode string `json:zip_code`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{Name: "Osawa Koki", City: "Soka", Zipcode: "340-0021"},
		{Name: "Sakura Mana", City: "Shibuya", Zipcode: "150-0043"},
		{Name: "Matz", City: "Tsukuba", Zipcode: "305-8577"},
	}

	// レスポンスヘッダにコンテントタイプを指定
	w.Header().Add("Content-Type", "application/json")

	// JSON形式にエンコード
	// IOライターを受け取る。(FPrint関数的な、、、)
	json.NewEncoder(w).Encode(customers)
}

func main() {

	// ハンドラファンクション -> デフォルトmultiplexerに登録。
	// 第一引数 -> パターン
	// 第二引数 -> 「レスポンスライター」「リクエスト」を引数として受け取る関数
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// 第一引数 -> リッスンするアドレス
	// 第二引数 -> 使用するmultiplexer | デフォルトで用意されているものを使用するため、nilを設定
	http.ListenAndServe("localhost:8000", nil)
}
