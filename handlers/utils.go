package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fredele20/rssagg/internal/database"
)

type ApiConfig struct {
	DB *database.Queries
}

func handlerErr(w http.ResponseWriter, r *http.Request) {
	responseWithError(w, 400, "something went wrong")
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5XX error: ", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{})
}
