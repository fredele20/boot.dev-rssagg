package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fredele20/rssagg/internal/auth"
	"github.com/fredele20/rssagg/models"
)

func (c *Core) CreateFeed(w http.ResponseWriter, r *http.Request) {
	// type parameter struct {
	// 	Name string `json:"name"`
	// 	URL  string `json:"url"`
	// }

	decoder := json.NewDecoder(r.Body)

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := c.core.GetUser(r.Context(), apiKey)

	var payload models.Feed
	err = decoder.Decode(&payload)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	payload.UserID = user.ID

	feed, err := c.core.CreateFeed(r.Context(), payload)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't create a feed: %v", err))
		return
	}

	respondWithJSON(w, 201, feed)
}

func (c *Core) GetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := c.core.GetFeeds(r.Context())
	if err != nil {
		responseWithError(w, 500, fmt.Sprintf("could not get feeds"))
		return
	}

	respondWithJSON(w, 200, feeds)
}
