package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fredele20/rssagg/internal/database"
	"github.com/google/uuid"
)


func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameter struct {
		Name string `json:"name"`
		URL string `json:"url"`
	}

	decoder := json.NewDecoder(r.Body)
	
	params := parameter{}
	err := decoder.Decode(&params)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
		Url: params.URL,
		UserID: user.ID,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't create a feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("could not get feeds"))
		return
	}

	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
}