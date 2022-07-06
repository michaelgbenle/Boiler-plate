package router

import (
	"github.com/gin-gonic/gin"
	"github.com/michaelgbenle/Boiler-plate/handlers"
	"os"
)

func SetupRouter(h *handlers.Handler) (*gin.Engine, string) {
	router := gin.Default()

	port := os.Getenv("PORT")

	return router, port
}
