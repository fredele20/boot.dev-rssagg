package handlers

import "github.com/go-chi/chi"


func RegisterRoutes(route *chi.Mux, core Core) {
	route.Get("/healthz", handlerReadiness)
	route.Get("/err", handlerErr)
	route.Post("/users", core.CreateUser)
	route.Get("/users", core.GetUser)
	// route.Get("/users", apiCfg.middlewareAuth(apiCfg.GetUser))

	route.Post("/feeds", core.CreateFeed)
	route.Get("/feeds", core.GetFeeds)

	route.Get("/posts", core.GetPostsForUser)


	route.Post("/feed_follows", core.CreateFeedFollow)
	route.Get("/feed_follows", core.GetFeedFollows)
	route.Delete("/feed_follows/{feedFollowID}", core.DeleteFeedFollow)
}

	