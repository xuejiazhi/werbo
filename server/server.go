package server

import (
	"io"
	"log"
	"net/http"
	"strings"
	"werbo/cgi"
	"werbo/configure"
)

func init() {

}

type Server struct{
	Port  string
	Resp http.ResponseWriter
	Req *http.Request
}

//CreateServer of go
func CreateServer() *Server {
	//get port
	Port,_ := configure.LoadCfg("werbo", "server","port")
	
	//listen port
	errCode := http.ListenAndServe(":" + Port, nil)
	if errCode != nil {
		log.Fatal("ListenAndServe: ", errCode)
	}

	//return
	retServer := new(Server)
	retServer.Port = Port
	return retServer
}//end func CreateServer


//GET Handler
func (this *Server) IndexHandler(rp http.ResponseWriter, rq *http.Request) {
	//Get URI
	addrList := strings.Split(this.Req.RequestURI, "/")
	
	var ret string
	//judgement params
	if len(addrList) >= 3 && addrList[1] != "" && addrList[2] != "" {
		dir := string(addrList[1])
		method := strings.TrimSpace(addrList[2])
		ret = cgi.Say(dir, method, this.Req, this.Resp)
	} else {
		ret = "welcome use werbo go"
	}
	//write string
	io.WriteString(this.Resp, ret)
}//end func IndexHandler
