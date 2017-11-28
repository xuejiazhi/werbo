package main

import(
	// "fmt"
	//  "werbo/configure"
	"net/http"
	 "werbo/server"
)
func main() {
	//创建一个httpserver
	c:=server.CreateServer()
	http.HandleFunc("/", c.IndexHandler) 
}

// func main() {
// 	fmt.Println(configure.LoadCfg("werbo", "base","date"))
// }
