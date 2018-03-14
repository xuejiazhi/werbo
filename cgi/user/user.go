package user

import "net/http"

type User struct {
}

func (this *User) Cgi_say(rq *http.Request, rp http.ResponseWriter) string {
	return "hahaha"
}
