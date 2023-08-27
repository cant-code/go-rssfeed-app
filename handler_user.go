package main

import (
	"encoding/json"
	"fmt"
	"github.com/cant-code/go-rssfeed-app/internal/database"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldnt create user: %v", err))
		return
	}

	respondWithJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) getUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJson(w, 200, databaseUserToUser(user))
}
