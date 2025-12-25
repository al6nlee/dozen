package initialize

import (
	"context"
	"dozen/backend/internal/global"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitRedis() {
	cfg := global.ServerConfig.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
		// ===== 连接池 =====
		PoolSize:     50,              // 最大连接数
		MinIdleConns: 10,              // 最小空闲连接
		PoolTimeout:  5 * time.Second, // 获取连接超时

		// ===== 超时控制 =====
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,

		// 空闲连接关闭时间
		ConnMaxIdleTime: 10 * time.Minute,
	})
	// 启动时校验
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Errorf("redis ping failed: %w", err))
	}
	global.Redis = rdb
}

func InitPostgres() {
	cfg := global.ServerConfig.Postgres
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
		cfg.Port,
	)

	logLevel := logger.Warn
	if global.ServerConfig.Mode == "dev" {
		logLevel = logger.Info
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logLevel,    // Log level
			Colorful:      false,       // 禁用彩色打印
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(fmt.Errorf("init postgres failed: %w", err))
	}

	// 获取底层 sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("get sql.DB failed: %w", err))
	}
	sqlDB.SetMaxOpenConns(50)                  // 最大打开连接数
	sqlDB.SetMaxIdleConns(10)                  // 最大空闲连接
	sqlDB.SetConnMaxLifetime(30 * time.Minute) // 连接最大生命周期, 必须设置，否则 PG 重启后容易拿到“僵尸连接”
	sqlDB.SetConnMaxIdleTime(10 * time.Minute) // 空闲连接最大时间
	if err := sqlDB.Ping(); err != nil {
		panic(fmt.Errorf("postgres ping failed: %w", err))
	}
	global.PG = db
}

func InitDB() {
	InitPostgres()
	InitRedis()
}
