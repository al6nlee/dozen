package initialize

import (
	"dozen/backend/internal/middleware"
	"dozen/backend/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func registerAuthRouter(r *gin.RouterGroup) {
	authRouter := r.Group("/auth")
	{
		authRouter.POST("/register", utils.NotImplemented)
		authRouter.POST("/register/code", utils.NotImplemented)
		authRouter.POST("/login", utils.NotImplemented)
		authRouter.POST("/logout", utils.NotImplemented)
	}
}

func registerRouters(r *gin.RouterGroup) {
	registerAuthRouter(r)
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
