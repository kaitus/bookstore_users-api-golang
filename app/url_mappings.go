package app

import (
	"github.com/kaitus/bookstore_users-api-golang/controllers/ping"
	"github.com/kaitus/bookstore_users-api-golang/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
