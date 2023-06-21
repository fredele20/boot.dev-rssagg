package handlers

import "github.com/go-chi/chi"


func RegisterRoutes(route *chi.Mux, core Core) {
	route.Get("/healthz", handlerReadiness)
	route.Get("/err", handlerErr)
	route.Post("/users", core.CreateUser)
	route.Get("/users", core.GetUser)
	// route.Get("/users", apiCfg.middlewareAuth(apiCfg.GetUser))

	// route.Post("/feeds", apiCfg.middlewareAuth(apiCfg.CreateFeed))
	// route.Get("/feeds", apiCfg.GetFeeds)

	// route.Get("/posts", apiCfg.middlewareAuth(apiCfg.GetPostsForUser))


	// route.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.CreateFeedFollow))
	// route.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.GetFeedFollows))
	// route.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.DeleteFeedFollow))
}

	