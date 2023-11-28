package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerServeSignupPage(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := getSessionToken(cfg.sessionSecret, r)
	if err != nil || sessionToken == nil {
		respondWithErrorPage(w, http.StatusUnauthorized, "Your session is invalid, hit the button below to head back home")
		return
	}
	userDataString, err := sessionToken.Claims.GetSubject()
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	userTwitterUsername := strings.Split(userDataString, " ")[1]
	userTwitterID := strings.Split(userDataString, " ")[0]
	_, err = cfg.DB.GetApplicantByTwitterID(r.Context(), userTwitterID)
	if err == nil {
		http.Redirect(w, r, "/waitlist/profile", http.StatusFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(signupPageHTMLString, userTwitterUsername)))
}
