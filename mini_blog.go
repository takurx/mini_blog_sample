package main

import (
	"github.com/bmizerany/pat"
	"html/template"
	"log"
	"net/http"
)

const (
	PORT = ":3000"
)

type Article struct {
	Index   int
	Title   string
	Content string
}

func main() {
	// ルーティング
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(indexHandler))
	mux.Get("/article/:url", http.HandlerFunc(articleHandler))
	mux.Post("/preview", http.HandlerFunc(previewHandler))
	mux.Post("/post", http.HandlerFunc(postHandler))
	http.Handle("/", mux)
	// サーバの起動 エラー吐いたらログ吐いて終了
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// トップページ
func indexHandler(w http.ResponseWriter, r *http.Request) {
	index, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "SERVER ERROR", 500)
	}
	index.Execute(w, nil)
}

// 記事ページ
func articleHandler(w http.ResponseWriter, r *http.Request) {
	// 以下の記述でURLのアレをアレできる 保存したデータの取り出しに使おう
	// 例: /article/aiueo → "aiueo"
	/*
		params := r.URL.Query()
		gameName := params.Get(":url")
	*/
}

// プレビューページ
func previewHandler(w http.ResponseWriter, r *http.Request) {
	pre, err := template.ParseFiles("article.html")
	if err != nil {
		http.Error(w, "SERVER ERROR", 500)
	}
	data := &Article{
		Title:   r.FormValue("title"),
		Content: r.FormValue("text"),
	}
	pre.Execute(w, data)
}

// 投稿を受け付けるハンドラ
func postHandler(w http.ResponseWriter, r *http.Request) {
	/*
		本来は r.Methodが"POST"なのか確認する必要があるが、今回はpatがPOSTを保証してくれている
		data := &Article{
			Index: r.FormValue("url"),
			Title: r.FormValue("title"),
			Text:  r.FormValue("text"),
		}
	*/
	// FormValueの漏れが無いか確認する必要もあるよ(特にurl)
}
