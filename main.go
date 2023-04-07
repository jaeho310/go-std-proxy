package main

import (
	"fmt"
	"go-std-proxy/server"
)

func main() {
	fmt.Println("proxy-server start")
	server.InitProxyServer()
}
