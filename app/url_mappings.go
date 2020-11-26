package app

import (
	"github.com/kaitus/bookstore_users-api-golang/controllers/ping"
	"github.com/kaitus/bookstore_users-api-golang/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/seach", users.Search)
}
