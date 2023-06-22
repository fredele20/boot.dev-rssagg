package core

import (
	"context"
	"time"

	"github.com/fredele20/rssagg/internal/database"
	"github.com/fredele20/rssagg/models"
	"github.com/google/uuid"
)

// func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
// 	type parameter struct {
// 		Name string `json:"name"`
// 		URL string `json:"url"`
// 	}

// 	decoder := json.NewDecoder(r.Body)

// 	params := parameter{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
// 		return
// 	}

// 	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
// 		ID: uuid.New(),
// 		CreatedAt: time.Now().UTC(),
// 		UpdatedAt: time.Now().UTC(),
// 		Name: params.Name,
// 		Url: params.URL,
// 		UserID: user.ID,
// 	})

// 	if err != nil {
// 		responseWithError(w, 400, fmt.Sprintf("couldn't create a feed: %v", err))
// 		return
// 	}

// 	respondWithJSON(w, 201, databaseFeedToFeed(feed))
// }

// func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
// 	feeds, err := apiCfg.DB.GetFeeds(r.Context())
// 	if err != nil {
// 		responseWithError(w, 500, fmt.Sprintf("could not get feeds"))
// 		return
// 	}

// 	respondWithJSON(w, 200, databaseFeedsToFeeds(feeds))
// }

func (a *ApiConfig) CreateFeed(context context.Context, payload models.Feed) (*models.Feed, error) {

	feed, err := a.DB.CreateFeed(context, database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      payload.Name,
		Url:       payload.Url,
	})
	if err != nil {
		return nil, err
	}

	newFeed := models.DatabaseFeedToFeed(feed)

	return &newFeed, nil
}

func (a *ApiConfig) GetFeeds(context context.Context) (*[]models.Feed, error) {
	feeds, err := a.DB.GetFeeds(context)
	if err != nil {
		return nil, err
	}

	dbFeeds := models.DatabaseFeedsToFeeds(feeds)

	return &dbFeeds, nil
}
