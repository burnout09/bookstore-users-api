package app

import "github.com/burnout09/bookstore-users-api/controllers"

func MapUrls() {
	router.GET("/ping", controllers.Ping)
	router.POST("/users/:user_id", controllers.GetUser)
	router.POST("/users/search", controllers.SearchUser)
	router.POST("/users", controllers.CreateUser)
}
