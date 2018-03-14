package cgi

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"werbo/cgi/test"
	"werbo/cgi/user"
)

//路由
var routerList = map[string]interface{}{
	"user": &user.User{},
	"test": &test.Test{},
}

func Say(router, method string, rq *http.Request, rp http.ResponseWriter) string {
	s := routerList[router]
	v := reflect.ValueOf(s)
	param := make([]reflect.Value, 2)
	param[0] = reflect.ValueOf(rq)
	param[1] = reflect.ValueOf(rp)
	method = "Cgi_" + strings.ToLower(method)
	fmt.Println("method====", method)
	m := v.MethodByName(method)
	fmt.Println("a====", m)
	//是否是正确的value
	if m.IsValid() {
		resp := m.Call(param)
		return resp[0].String()
	} else {
		return "方法错误"
	}
}
