package server

import (
	"io"
	"log"
	"net/http"
	"strings"
	"werbo/cgi"
)

func init() {

}

func CreateServer() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func IndexHandler(rp http.ResponseWriter, rq *http.Request) {
	//获取URI
	uriList := strings.Split(rq.RequestURI, "/")
	var retStr string
	if len(uriList) >= 3 && uriList[1] != "" && uriList[2] != "" {
		dir := string(uriList[1])
		method := strings.TrimSpace(uriList[2])
		retStr = cgi.Say(dir, method, rq, rp)
	} else {
		retStr = "welcome use jee go"
	}
	//写入
	io.WriteString(rp, retStr)
}
