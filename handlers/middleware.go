package handlers

import (
	"fmt"
	"net/http"

	"github.com/fredele20/rssagg/internal/auth"
	"github.com/fredele20/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *Core) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := cfg.core.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, 500, fmt.Sprintf("couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
