package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fredele20/rssagg/core"
	"github.com/fredele20/rssagg/internal/auth"
	"github.com/fredele20/rssagg/internal/database"
	"github.com/fredele20/rssagg/models"
)

type Core struct {
	core *core.ApiConfig
}

func NewCore(c *core.ApiConfig) *Core {
	return &Core{
		core: c,
	}
}

func (c *Core) CreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	var payload models.User
	err := decoder.Decode(&payload)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := c.core.CreateUser(r.Context(), payload)

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
	}

	respondWithJSON(w, 201, user)
}

func (c *Core) GetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}
	
	fmt.Println(apiKey)
	user, err := c.core.GetUser(r.Context(), apiKey)
	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("could not get user: %v", err))
		return
	}

	respondWithJSON(w, 200, user)
}

func (apiCfg *ApiConfig) GetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, models.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) GetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  10,
	})

	if err != nil {
		responseWithError(w, 400, fmt.Sprintf("could not get posts: %v", err))
		return
	}

	respondWithJSON(w, 200, models.DatabasePostsToPosts(posts))
}
