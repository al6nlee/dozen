package main

import (
	"context"
	"dozen/backend/internal/global"
	"dozen/backend/internal/initialize"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CloseDB() {
	if global.PG != nil {
		sqlDB, _ := global.PG.DB()
		_ = sqlDB.Close()
	}
	if global.Redis != nil {
		_ = global.Redis.Close()
	}
}

func main() {
	initialize.InitConfig()
	initialize.InitLogger()

	zap.S().Infof("DoZen Backend Server Starting...")
	initialize.InitDB()

	if global.ServerConfig.Mode == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := initialize.InitRouters()
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", global.ServerConfig.PORT),
		Handler: router,
	}

	// Start HTTP server
	go func() {
		zap.S().Infof("HTTP server listening on :%d", global.ServerConfig.PORT)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.S().Fatal("HTTP server start failed", err)
		}
	}()

	// 监听系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.S().Info("Shutting down server...")

	// 优雅关闭
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.S().Error("Server shutdown failed", err)
	}
	CloseDB()
	zap.S().Info("DoZen Backend Server exited")
}
