package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerLogout(w http.ResponseWriter, r *http.Request) {
	a := destroyCookie("sessionID")
	b := destroyCookie("waitlist-referrer")
	w.Header().Set("Set-Cookie", a.String())
	w.Header().Add("Set-Cookie", b.String())
	http.Redirect(w, r, "/waitlist/", http.StatusFound)
}
