package test

import (
	"net/http"
)

type Test struct {
}

func (this *Test) Getuserifo(rq *http.Request, rp http.ResponseWriter) string {
	return `{"result":-1,"msg":"faild to output json"}`
}

func (this *Test) Sayok(rq *http.Request, rp http.ResponseWriter) string {
	return `{"result":0,"msg":"ok"}`
}
