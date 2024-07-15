package todos

import (
	"main/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 全件取得エンドポイントのハンドラ
func GetAll(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var todos []models.Todo

	err := db.Find(&todos).Error

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
	} else {
		ctx.JSON(http.StatusOK, todos)
	}
}

// ID取得エンドポイントのハンドラ
func GetById(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var todo models.Todo
	err := db.First(&todo, ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}
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
	db := ctx.MustGet("db").(*gorm.DB)

	var requestBody models.Todo
	err := ctx.BindJSON(&requestBody)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "不正なリクエストボディです",
		})
	} else {
		var old models.Todo

		err := db.First(&old, ctx.Param("id")).Error
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "not found",
			})
		}
		requestBody.Id = old.Id
		db.Save(&requestBody)

		ctx.JSON(http.StatusOK, requestBody)
	}
}

// 削除エンドポイントのハンドラ
func DeleteById(ctx *gin.Context) {
	db := ctx.MustGet("db").(*gorm.DB)
	var old models.Todo
	err := db.First(&old, ctx.Param("id")).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "not found",
		})
	}

	if err := db.Delete(&old).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
	} else {
		ctx.Status(http.StatusNoContent)
	}
}

func RegisterTodoEndpoints(todoGroup *gin.RouterGroup) {
	todoGroup.GET("", GetAll)
	todoGroup.GET(":id", GetById)
	todoGroup.POST("", CreateNew)
	todoGroup.DELETE(":id", DeleteById)
	todoGroup.PUT(":id", Update)
}
