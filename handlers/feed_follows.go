package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fredele20/rssagg/internal/auth"
	"github.com/fredele20/rssagg/models"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (c *Core) CreateFeedFollow(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)

	var payload models.FeedFollow
	err := decoder.Decode(&payload)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := c.core.GetUser(r.Context(), apiKey)

	payload.UserID = user.ID

	feedFollow, err := c.core.CreateFeedFollow(r.Context(), payload)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, feedFollow)
}

func (c *Core) GetFeedFollows(w http.ResponseWriter, r *http.Request) {

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := c.core.GetUser(r.Context(), apiKey)

	feedFollows, err := c.core.GetFeedFollows(r.Context(), user.ID)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't get feed follows: %v", err))
		return
	}

	respondWithJSON(w, 200, feedFollows)
}

func (c *Core) DeleteFeedFollow(w http.ResponseWriter, r *http.Request) {
	feedFollowIDStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't parse feed follow id:%v", err))
		return
	}

	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	user, err := c.core.GetUser(r.Context(), apiKey)

	err = c.core.DeleteFeedFollow(r.Context(), feedFollowID, user.ID)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("couldn't delete feed follow: %v", err))
		return
	}

	respondWithJSON(w, 200, struct{}{})
}
