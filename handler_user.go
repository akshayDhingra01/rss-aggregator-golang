package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/akshayDhingra01/rss-aggregator-golang/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) { // funct signature can't be changed in go but we can create this as an method and pass api config

	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Passing json: %s", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		ModeifiedAt: time.Now().UTC(),
		Name:        params.Name,
	},
	)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Could not create user: %s", err))
		return
	}

	respondWithJson(w, 200, user)

}
