package httpapi

import (
	"github.com/gin-gonic/gin"

	"github.com/berjis/markets/backend/internal/httpapi/routes"
)

// Register attaches all HTTP routes to the router.
func Register(router *gin.Engine) {
	api := router.Group("/api")
	routes.RegisterHealth(api)
}
