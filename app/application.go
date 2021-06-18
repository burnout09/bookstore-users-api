package app

import (
	"github.com/burnout09/bookstore-users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	MapUrls()

	logger.Info("about to start the application...")
	err := router.Run(":8080")
	if err != nil {
		return 
	}
}
