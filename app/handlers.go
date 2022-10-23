package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

// JSON解析用に使用する構造体の定義
// JSON出力用のキーは「`json:"★★★"`」と指定することが可能。
// 指定しなければ、構造体のキーがそのまま使用される。
// XMLを使用する場合には「xml:"★★★"」を書く。
// ひとつの構造体をJSONとXMLで共有することも可能。
type Customer struct {
	Name    string `json:"fill_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
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

func getAllCustomers_xml(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Osawa Koki", City: "Soka", Zipcode: "340-0021"},
		{Name: "Sakura Mana", City: "Shibuya", Zipcode: "150-0043"},
		{Name: "Matz", City: "Tsukuba", Zipcode: "305-8577"},
	}

	// レスポンスヘッダにコンテントタイプを指定
	w.Header().Add("Content-Type", "application/xml")

	// XML形式にエンコード
	xml.NewEncoder(w).Encode(customers)
}

func getAllCustomers_flex(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Osawa Koki", City: "Soka", Zipcode: "340-0021"},
		{Name: "Sakura Mana", City: "Shibuya", Zipcode: "150-0043"},
		{Name: "Matz", City: "Tsukuba", Zipcode: "305-8577"},
	}

	if r.Header.Get("Accept") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
