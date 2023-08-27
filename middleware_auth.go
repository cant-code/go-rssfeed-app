package main

import (
	"fmt"
	"github.com/cant-code/go-rssfeed-app/internal/auth"
	"github.com/cant-code/go-rssfeed-app/internal/database"
	"net/http"
)

type authedHandler func(w http.ResponseWriter, response *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), key)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Couldnt find user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
