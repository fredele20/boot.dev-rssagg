package core

import (
	"context"
	"time"

	"github.com/fredele20/rssagg/internal/database"
	"github.com/fredele20/rssagg/models"
	"github.com/google/uuid"
)

func (a *ApiConfig) CreateFeedFollow(context context.Context, payload models.FeedFollow) (*models.FeedFollow, error) {

	feedFollow, err := a.DB.CreateFeedFollow(context, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    payload.UserID,
		FeedID:    payload.FeedID,
	})

	if err != nil {
		return nil, err
	}

	feedFollowDB := models.DatabaseFeedFollowToFeedFollow(feedFollow)

	return &feedFollowDB, nil
}

func (a *ApiConfig) GetFeedFollows(context context.Context, userId uuid.UUID) (*[]models.FeedFollow, error) {

	feedFollows, err := a.DB.GetFeedFollows(context, userId)
	if err != nil {
		return nil, err
	}

	feedFollowsDB := models.DatabaseFeedFollowsToFeedFollows(feedFollows)

	return &feedFollowsDB, nil
}

func (a *ApiConfig) DeleteFeedFollow(context context.Context, id, userId uuid.UUID) error {

	err := a.DB.DeleteFeedFollow(context, database.DeleteFeedFollowParams{
		ID:     id,
		UserID: userId,
	})

	if err != nil {
		return err
	}

	return nil
}
