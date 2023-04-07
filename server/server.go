package server

import (
	"go-proxy/handler"
	"net/http"
)

func InitProxyServer() {
	mux := http.NewServeMux()
	handler.InitDefaultHandler(mux)
	handler.InitProxyHandler(mux)
	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}
