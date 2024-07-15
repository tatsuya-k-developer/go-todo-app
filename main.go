package main

import (
	"fmt"
	"main/config"
	"main/models"
	"main/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	}
}

func main() {
	cfg, err := config.Load("config.yaml")

	if err != nil {
		fmt.Println("設定ファイルを読み込めませんでした")
		return
	}

	db, err := gorm.Open(postgres.Open(cfg.GetDBDNS()))

	if err != nil {
		fmt.Println("DBに接続できませんでした")
		return
	}

	// 自動でDBのマイグレーションを行ってくれる
	db.AutoMigrate(&models.Todo{})

	router := gin.Default()

	router.Use(DBMiddleware(db))

	routes.AddRoutes(router)

	router.Run(cfg.GetSocketAddr())
}
