package main

import (
	"database/sql"
	"encoding/json"
	"internal/database"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

type signupReqBody struct {
	Email         string `json:"email"`
	WalletAddress string `json:"wallet_address"`
}

func (cfg *apiConfig) handlerWaitlistSignup(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := getSessionToken(cfg.sessionSecret, r)
	if err != nil || sessionToken == nil {
		respondWithErrorPage(w, http.StatusUnauthorized, "Your session is invalid, hit the button below to head back home")
		log.Println(err.Error())
		return
	}
	userDataString, err := sessionToken.Claims.GetSubject()
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	bodyStruct := signupReqBody{}
	err = json.Unmarshal(body, &bodyStruct)
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	if len(bodyStruct.WalletAddress) != 42 || bodyStruct.WalletAddress[0:2] != "0x" {
		respondWithErrorPage(w, http.StatusInternalServerError, "Invalid wallet address")
		log.Println(err.Error())
		return
	}
	userData := strings.Split(userDataString, " ")
	twitterID := userData[0]
	twitterUsername := userData[1]
	applicantUUID, err := uuid.NewRandom()
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	var referrer sql.NullString
	cookie, err := r.Cookie("waitlist-referrer")
	if err == nil && cookie.Value != "" {
		_, err = cfg.DB.GetApplicantByTwitterID(r.Context(), cookie.Value)
		if err == nil {
			referrer = sql.NullString{
				String: cookie.Value,
				Valid:  true,
			}
		}
	}
	_, err = cfg.DB.AddToWaitlist(r.Context(), database.AddToWaitlistParams{
		ID:              applicantUUID,
		TwitterID:       twitterID,
		TwitterUsername: twitterUsername,
		Referrer:        referrer,
		Email:           bodyStruct.Email,
		WalletAddress:   bodyStruct.WalletAddress,
	})
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	http.Redirect(w, r, "/waitlist/profile", http.StatusFound)
}
