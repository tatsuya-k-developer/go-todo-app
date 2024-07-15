package routes

import (
	"main/routes/todos"
	
	"github.com/gin-gonic/gin"
)

// メイン側から呼び出すエンドポイントの登録のための関数
func AddRoutes(e *gin.Engine) {
	todoGroup := e.Group("/todo")
	todos.RegisterTodoEndpoints(todoGroup)
}