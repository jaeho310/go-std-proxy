package main

import (
	"fmt"
	"go-proxy/server"
)

func main() {
	fmt.Println("proxy-server start")
	server.InitProxyServer()
}
