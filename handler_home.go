package main

import (
	"encoding/base64"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerHome(w http.ResponseWriter, r *http.Request) {
	ok, _ := hasValidSessionToken(cfg.sessionSecret, r)
	if !ok {
		refEncoded := r.URL.Query().Get("ref")
		var referrer string
		if refEncoded != "" {
			decoded, err := base64.URLEncoding.DecodeString(refEncoded)
			if err != nil {
				respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
				log.Println(err.Error())
				return
			}
			referrer = string(decoded)
		}
		cookie := http.Cookie{
			Name:     "waitlist-referrer",
			Value:    referrer,
			MaxAge:   604800,
			Domain:   "localhost",
			Path:     "/",
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		}
		w.Header().Set("Set-Cookie", cookie.String())
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(homePageHTMLString))
		return
	}
	sessionToken, err := getSessionToken(cfg.sessionSecret, r)
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	sessionUsername, err := sessionToken.Claims.GetSubject()
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	userTwitterID := strings.Split(sessionUsername, " ")[0]
	_, err = cfg.DB.GetApplicantByTwitterID(r.Context(), userTwitterID)
	if err != nil {
		http.Redirect(w, r, "/waitlist/signup", http.StatusFound)
		return
	}
	http.Redirect(w, r, "/waitlist/profile", http.StatusFound)
}
