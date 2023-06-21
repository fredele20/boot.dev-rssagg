package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fredele20/rssagg/core"
	"github.com/fredele20/rssagg/handlers"
	"github.com/fredele20/rssagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

// type coreConfig struct {
// 	c *core.ApiConfig
// }

func main() {

	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can not connect to database:", err)
	}

	db := database.New(conn)
	apiCfg := &apiConfig{
		DB: db,
	}

	handlersService := handlers.NewCore((*core.ApiConfig)(apiCfg))

	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// routes := *chi.NewMux()

	handlers.RegisterRoutes(router, *handlersService)

	// v1Router.Get("/healthz", handlerReadiness)
	// v1Router.Get("/err", handlerErr)
	// v1Router.Post("/users", apiCfg.handlerCreateUser)
	// v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	// v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	// v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	// v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	// v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	// v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	// v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	// router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + port,
	}

	log.Printf("Server started on port %v", port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
