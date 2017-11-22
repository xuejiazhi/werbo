package main

import "werbo/utils"

// func main() {
// 	//创建一个httpserver
// 	http.HandleFunc("/", server.IndexHandler)
// 	server.CreateServer()
// }

func main() {
	utils.LoadConfig("werbo", "base")
}
