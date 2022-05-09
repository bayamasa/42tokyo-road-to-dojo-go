package main

import (
	"flag"

	"42tokyo-road-to-dojo-go/pkg/infra/http"
)

var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	http.Serve(addr)
}
