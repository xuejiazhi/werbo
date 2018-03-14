package test

import (
	"net/http"
)

type Test struct {
}

func (this *Test) Cgi_getuserinfo(rq *http.Request, rp http.ResponseWriter) string {
	return `{"result":-1,"msg":"faild to output json"}`
}

func (this *Test) Cgi_sayok(rq *http.Request, rp http.ResponseWriter) string {
	return `{"result":0,"msg":"ok"}`
}
