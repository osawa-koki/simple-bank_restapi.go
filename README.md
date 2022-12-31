# simple-restapi_outline.go

RESTfulアーキテクチャの基づいたGo言語による学習目的APIサーバ構築サンプル。  

## 実行方法

```shell
# デバグ実行
go run main.go

# ビルド
go build -o bin main.go
```

## モジュール管理

同一ディレクトリに存在するファイルでひとつのモジュールを構築する。  
したがって、他の言語と異なり、サブディレクトリにあるファイルは別のファイルに属する。  

モジュールは当該ディレクトリにある`go.mod`ファイルで管理される。  
`module example.com`は当該モジュール名が`example.com`であることを意味する。  

モジュール名は任意で構わないが、実在する他のモジュールを指定することはできない。  

`example.com`モジュール内のファイルで`import "example.com/app"`と書くと`example.com`モジュールのサブディレクトリである`app`ディレクトリから構成されるモジュールを使用可能になる。  

このプロジェクトではルートリポジトリに`example.com`という名前のモジュールを作成し、`main.go`ファイル内で`import "example.com/app"`と書いているため、appディレクトリにある`app.go`と`handlers.go`から構成されるモジュールを使用可能。  

## 実行方法(Docker)

```shell
# Dockerfileのビルドの実行
docker build -t simple-restapi_outline.go .
docker run -p 80:80 -it --rm --name my-simple-restapi_outline.go simple-restapi_outline.go

# 一行で書くなら、、、
docker build -t simple-restapi_outline.go . && docker run -p 80:80 -it --rm --name my-simple-restapi_outline.go simple-restapi_outline.go
```

## 動作確認

以下のパスに対してアクセスする。  

- /
- /greet
- /customers
- /customers_xml
- /customers_flex

また、それぞれ`Accept`ヘッダを`text/html`と`application/json`にセットして試してみる。  

```go
type Customer struct {
  Name    string `json:"fill_name" xml:"name"`
  City    string `json:"city" xml:"city"`
  Zipcode string `json:"zip_code" xml:"zipcode"`
}
```

json用とxml用のプロパティは構造体で設定し、それぞれの出し分けは`r.Header.Get("Accept") == "text/xml"`で判断している。  

## デプロイ設定(Render.com)

| キー | バリュー |
| ---- | ---- |
| Name | simple-restapi_outline.go |
| Region | Oregon(US West) |
| Branch | main |
| Root Directory |  |
| Environment | Docker |
| Dockerfile Path | ./Dockerfile |
| Docker Build Context Directory |  |
| Docker Command |  |

## 参考文献

- <https://www.udemy.com/share/103PE43@JyZWqaqer4WCJGGiYgQalepcfxVW38cHTmN7p5IyKIbP0xHR4WVlUMYrBZDi_oGr8g==/>
