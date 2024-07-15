package todos

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 全件取得エンドポイントのハンドラ
func GetAll(ctx *gin.Context) {

}

// ID取得エンドポイントのハンドラ
func GetById(ctx *gin.Context) {

}

// 新規作成エンドポイントのハンドラ
func CreateNew(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)

	var requestBody models.Todo

	err := ctx.BindJSON(&requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "不正なリクエストボディです",
		})
	} else {
		// 作成処理を行う
		// DBにデータをインサートしないといけない
		db.Create(&requestBody)
		ctx.JSON(http.StatusCreated, requestBody)
	}
}

// 更新エンドポイントのハンドラ
func Update(ctx *gin.Context) {

}

// 削除エンドポイントのハンドラ
func DeleteById(ctx *gin.Context) {

}

func RegisterTodoEndpoints(todoGroup *gin.RouterGroup) {
	todoGroup.GET("", GetAll)
	todoGroup.GET(":id", GetById)
	todoGroup.POST("", CreateNew)
	todoGroup.DELETE(":id", DeleteById)
	todoGroup.PUT(":id", Update)
}
