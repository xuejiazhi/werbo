package main

import(
	"fmt"
	 "werbo/configure"
)
// func main() {
// 	//创建一个httpserver
// 	http.HandleFunc("/", server.IndexHandler)
// 	server.CreateServer()
// }

func main() {
	fmt.Println(configure.LoadCfg("werbo", "base","date"))
}
