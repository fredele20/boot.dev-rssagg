package core

import (
	"context"
	"time"

	"github.com/fredele20/rssagg/internal/database"
	"github.com/fredele20/rssagg/models"
	"github.com/google/uuid"
)

type ApiConfig struct {
	DB *database.Queries
}

func NewApiConfig(db *database.Queries) *ApiConfig {
	return &ApiConfig{
		DB: db,
	}
}

func (a *ApiConfig) CreateUser(context context.Context, payload models.User) (*models.User, error) {
	// type parameters struct {
	// 	Name string `json:"name"`
	// }

	user, err := a.DB.CreateUser(context, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      payload.Name,
	})

	if err != nil {
		return nil, err
	}

	newUser := models.DatabaseUserToUser(user)

	return &newUser, nil
}

func (a *ApiConfig) GetUser(context context.Context, apiKey string) (*models.User, error) {
	user, err := a.DB.GetUserByAPIKey(context, apiKey)
	if err != nil {
		return nil, err
	}
	dbUser := models.DatabaseUserToUser(user)
	return &dbUser, nil
}
