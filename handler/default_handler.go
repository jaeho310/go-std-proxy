package handler

import (
	"fmt"
	"net/http"
)

func InitDefaultHandler(mux *http.ServeMux) {
	mux.HandleFunc("/health", health)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello~")
}
