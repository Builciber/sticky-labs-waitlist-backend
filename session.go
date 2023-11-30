package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dghubble/gologin/v2/twitter"
	"github.com/golang-jwt/jwt/v5"
)

func (cfg *apiConfig) issueSessionToken() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		twitterUser, err := twitter.UserFromContext(ctx)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Something went wrong -> %s", err))
			return
		}
		currentTimeInUTC := time.Now().UTC()
		expirationTimeinUTC := currentTimeInUTC.Add(24 * time.Hour)
		sessionClaims := &jwt.RegisteredClaims{
			Issuer:    "waitlist-session",
			IssuedAt:  jwt.NewNumericDate(currentTimeInUTC),
			ExpiresAt: jwt.NewNumericDate(expirationTimeinUTC),
			Subject:   fmt.Sprintf("%d %s", twitterUser.ID, twitterUser.ScreenName),
		}
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaims)
		signedSessionToken, err := accessToken.SignedString([]byte(cfg.sessionSecret))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error creating signed session token")
			return
		}
		sessionCookie := http.Cookie{
			Name:       "sessionID",
			Value:      signedSessionToken,
			Expires:    expirationTimeinUTC,
			Domain:     cfg.domain,
			Path:       "/",
			HttpOnly:   true,
			Secure:     false,
			SameSite:   http.SameSiteLaxMode,
			RawExpires: expirationTimeinUTC.String(),
		}
		w.Header().Add("Set-Cookie", sessionCookie.String())
		w.Header().Set("Content-Type", "text/html")
		_, err = cfg.DB.GetApplicantByTwitterID(r.Context(), twitterUser.IDStr)
		if err != nil {
			http.Redirect(w, r, "/waitlist/signup", http.StatusFound)
			return
		}
		http.Redirect(w, r, "/waitlist/profile", http.StatusFound)
	}

	return http.HandlerFunc(fn)
}

func getSessionToken(sessionSecret string, r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie("sessionID")
	if err != nil {
		return nil, nil
	} //cookie with that name was not found
	sessionTokenString := cookie.Value
	tokenClaims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(
		sessionTokenString,
		&tokenClaims,
		func(t *jwt.Token) (interface{}, error) { return []byte(sessionSecret), nil })
	if err != nil {
		return nil, err
	}
	issuer, err := token.Claims.GetIssuer()
	if err != nil {
		return nil, err
	}
	if issuer != "waitlist-session" {
		return nil, nil
	}
	return token, nil
}

func hasValidSessionToken(sessionSecret string, r *http.Request) (bool, error) {
	sessionToken, err := getSessionToken(sessionSecret, r)
	if err != nil {
		return false, err
	}
	if sessionToken == nil {
		return false, nil
	}
	return true, nil
}

func (cfg *apiConfig) handlerIssueSessionToken() http.Handler {
	//A handler to test that the issuing of session tokens works as designed
	fn := func(w http.ResponseWriter, r *http.Request) {
		currentTimeInUTC := time.Now().UTC()
		expirationTimeinUTC := currentTimeInUTC.Add(24 * time.Hour)
		sessionClaims := &jwt.RegisteredClaims{
			Issuer:    "waitlist-session",
			IssuedAt:  jwt.NewNumericDate(currentTimeInUTC),
			ExpiresAt: jwt.NewNumericDate(expirationTimeinUTC),
			Subject:   fmt.Sprintf("%d %s", 8, "jill"),
		}
		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, sessionClaims)
		signedSessionToken, err := accessToken.SignedString([]byte(cfg.sessionSecret))
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error creating signed session token")
			return
		}
		sessionCookie := http.Cookie{
			Name:       "sessionID",
			Value:      signedSessionToken,
			Expires:    expirationTimeinUTC,
			Domain:     cfg.domain,
			Path:       "/",
			HttpOnly:   true,
			Secure:     false,
			SameSite:   http.SameSiteLaxMode,
			RawExpires: expirationTimeinUTC.String(),
		}
		w.Header().Add("Set-Cookie", sessionCookie.String())
		println(sessionCookie.String())
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte(fmt.Sprintf(`<p>You are logged in %s!</p><form action="/waitlist-app/logout" method="post"><input type="submit" value="Logout"></form>`, "builciber")))
		//http.Redirect(w, r, "/waitlist-app/", http.StatusFound)
	}

	return http.HandlerFunc(fn)
}

func (cfg *apiconfig) destroyCookie(cookieName string) http.Cookie {
	cookie := http.Cookie{
		Name:     cookieName,
		MaxAge:   -1,
		Domain:   cfg.domain,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
	return cookie
}
