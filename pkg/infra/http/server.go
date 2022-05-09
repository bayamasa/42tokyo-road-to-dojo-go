package http

import (
	"log"
	"net/http"

	"42tokyo-road-to-dojo-go/pkg/presentation/handler"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {

	/* ===== URLマッピングを行う ===== */
	http.HandleFunc("/setting/get", Get(handler.HandleSettingGet()))

	// TODO: 認証を行うmiddlewareを実装する
	// middlewareは pkg/http/middleware パッケージを利用する
	// http.HandleFunc("/user/get",
	//   get(middleware.Authenticate(handler.HandleUserGet())))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
