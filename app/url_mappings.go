package app

import (
	"github.com/burnout09/bookstore-users-api/controllers/ping"
	"github.com/burnout09/bookstore-users-api/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:user_id", users.Get)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)
	router.POST("/users/login", users.Login)
}
