package global

import (
	"dozen/backend/config"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	PG           *gorm.DB
	Redis        *redis.Client
)
