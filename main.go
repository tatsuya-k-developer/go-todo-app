package main

import (
	"fmt"
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
	dsn := "user=postgres password=password port=5432 sslmode=disable host=127.0.0.1"

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		fmt.Println("DBに接続できませんでした")
		return
	}

	// Todoのテーブルを自動的に更新する　Migrate＝アップデート
	// 実開発ではあんまりしないほうがいい＝バージョン管理ができなくなる　Goのマイグレーションツールを使った方が良さそう
	db.AutoMigrate(&models.Todo{})

	router := gin.Default()

	router.Use(DBMiddleware(db))

	routes.AddRoutes(router)

	router.Run("127.0.0.1:8080")
}