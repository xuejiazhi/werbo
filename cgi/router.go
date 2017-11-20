package cgi

import (
	"net/http"
	"reflect"
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
	m := v.MethodByName(method)
	//是否是正确的value
	if m.IsValid() {
		resp := m.Call(param)
		return resp[0].String()
	} else {
		return "方法错误"
	}
}
