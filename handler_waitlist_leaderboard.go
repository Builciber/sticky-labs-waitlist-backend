package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (cfg *apiConfig) handlerWaitlistLeaderboard(w http.ResponseWriter, r *http.Request) {
	limitAsString := r.URL.Query().Get("limit")
	limit := 50
	var err error
	if limitAsString != "" {
		limit, err = strconv.Atoi(limitAsString)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error parsing limit parameter as int")
			return
		}
	}
	leaderboardMembers, err := cfg.DB.GetLeaderboard(r.Context(), int32(limit))
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Something went wrong -> %s", err))
		return
	}
	respondWithJSON(w, 200, leaderboardMembers)
}
