package user

import "net/http"

type User struct {
}

func (this *User) Say(rq *http.Request, rp http.ResponseWriter) string {
	return "hahaha"
}
