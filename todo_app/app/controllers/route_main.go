package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request){
	// ParseFiles関数を使って、テンプレートファイルをパースする
	// htmlテンプレートファイルをパースし、テンプレートオブジェクトが返される
	// テンプレートオブジェクトのExecuteメソッドを使って、テンプレートをブラウザにレンダリングする
	// t, err := template.ParseFiles("app/views/templates/top.html")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// 第二引数
	// Executeメソッドの第二引数には、テンプレートに埋め込むデータを渡す
	// ここでは、文字列"Hello, World"を渡している
	// テンプレートファイルには、{{.}}という記述があり、これはExecuteメソッドの第二引数で渡したデータを埋め込む場所を示している
	// err = t.Execute(w, "Hello")
	// if err != nil {
    //     log.Fatal(err)
    // }

	_, err := session(w, r)
	// sessionがなくてerrがnilでない場合、topページにリダイレクト
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	}else{
		// sessionがある場合、indexページを表示
		http.Redirect(w, r, "/todos", 302)
	}

}

// index関数を追加
func index(w http.ResponseWriter, r *http.Request){
	// session関数を使って、セッションを取得
	// session自体は使わないので、_で受け取る
	_, err := session(w, r)
	// sessionがなくてerrがnilでない場合、topページにリダイレクト
	if err != nil {
		http.Redirect(w, r, "/l", 302)
	}else{
		// sessionがある場合、indexページを表示
		generateHTML(w, nil, "layout", "private_navbar","index")
	}
}