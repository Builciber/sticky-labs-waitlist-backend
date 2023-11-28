package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerHomepageRedirection() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/waitlist/", http.StatusFound)
	}
	return http.HandlerFunc(fn)
}
