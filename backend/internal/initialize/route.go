package initialize

import (
	"dozen/backend/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerRouters(r *gin.RouterGroup) {
	// TODO: Register your routers here
}

func InitRouters() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.CORS())

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})
	rootRouter := engine.Group("/api/v1")
	registerRouters(rootRouter)
	return engine
}
