package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func (cfg *apiConfig) handlerServeProfilePage(w http.ResponseWriter, r *http.Request) {
	sessionToken, err := getSessionToken(cfg.sessionSecret, r)
	if err != nil || sessionToken == nil {
		respondWithErrorPage(w, http.StatusUnauthorized, "Your session is invalid, hit the button below to head back home")
		if err != nil {
			log.Println(err.Error())
		}
		return
	}
	userDataString, err := sessionToken.Claims.GetSubject()
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	userTwitterID := strings.Split(userDataString, " ")[0]
	refQueryValue := base64.URLEncoding.EncodeToString([]byte(userTwitterID))
	referralCount, err := cfg.DB.GetUserReferralCount(r.Context(), sql.NullString{
		String: userTwitterID,
		Valid:  true,
	})
	if err != nil {
		respondWithErrorPage(w, http.StatusInternalServerError, "Something went wrong, its not your fault")
		log.Println(err.Error())
		return
	}
	userReferralLink := fmt.Sprintf("https://www.stickylabs.xyz/waitlist/?ref=%s", refQueryValue)
	twitterPrefilledMessage := fmt.Sprintf("An%%20L2%%20but%%20for%%20NFTs%%2C%%20that%%27s%%20%%40Labs_Sticky.%%20Get%%20exclusive%%20access%%20to%%20Drops%%2C%%20%%25%%20in%%20revenue%%2C%%20pass%%20into%%20the%%20ecosystem%%20and%%20much%%20more.%%0A%%0ASign-up%%20for%%20wave%%202%%20%%F0%%9F%%91%%87%%3A%%0Ahttps%%3A%%2F%%2Fwww.stickylabs.xyz%%2Fwaitlist%%2F%%3Fref%%3D%s", refQueryValue)
	prefilledMessage := fmt.Sprintf("An%%20L2%%20but%%20for%%20NFTs%%2C%%20that%%27s%%20Sticky%%20Labs.%%20Get%%20exclusive%%20access%%20to%%20Drops%%2C%%20%%25%%20in%%20revenue%%2C%%20pass%%20into%%20the%%20ecosystem%%20and%%20much%%20more.%%0A%%0ASign-up%%20for%%20wave%%202%%20%%F0%%9F%%91%%87%%3A%%0Ahttps%%3A%%2F%%2Fwww.stickylabs.xyz%%2Fwaitlist%%2F%%3Fref%%3D%s", refQueryValue)
	twitterShareLink := fmt.Sprintf("https://twitter.com/intent/tweet?text=%s", twitterPrefilledMessage)
	redditShareLink := fmt.Sprintf("https://www.reddit.com/r/CryptoCurrency/submit?title=Sticky%%20Labs%%20Wave%%202&selftext=true&text=%s", prefilledMessage)
	telegramShareLink := fmt.Sprintf("https://t.me/share/?text=%s", prefilledMessage)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(profilePageHTMLString, twitterShareLink, redditShareLink, telegramShareLink, userReferralLink, referralCount)))
}
