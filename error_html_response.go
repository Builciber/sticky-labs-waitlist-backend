package main

import (
	"fmt"
	"net/http"
)

func respondWithErrorPage(w http.ResponseWriter, code int, errMsg string) {
	errPage := fmt.Sprintf(errorPageHTMLString, errMsg)
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(code)
	w.Write([]byte(errPage))
}
