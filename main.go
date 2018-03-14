package main

import (
	// "fmt"
	//  "werbo/configure"

	"net/http"
	"werbo/server"
)

func main() {
	//创建一个httpserver
	http.HandleFunc("/", server.IndexHandler)
	server.CreateServer()
}
