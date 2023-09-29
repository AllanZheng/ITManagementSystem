package api

import (
	"fmt"
	"log"
	"net/http"
	"system/model"

	"github.com/gin-gonic/gin"
)

// type Controller struct {
// 	DB *gorm.DB
// }

/*
创建账单记录
*/
func CreateBad(ctx *gin.Context) {

	var newRecord *model.FeeRecord
	// 数据验证
	if err := ctx.ShouldBind(&newRecord); err != nil || newRecord == nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"data": "数据验证错误",
			"msg":  err.Error(),
		})
		return
	}
	fmt.Println(newRecord)
	if err := model.Create(*newRecord); err != nil {
		log.Print(err.Error())
		return
	}

	// 成功
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": "",
		"msg":  "创建成功",
	})
}

/*
更新账单记录
*/
func UpdateBad(ctx *gin.Context) {

}

/*
删除账单记录
*/
func DeleteBad(ctx *gin.Context) {
	var req *model.Query
	// 数据验证
	if err := ctx.ShouldBind(&req); err != nil || req == nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"data": "数据验证错误",
			"msg":  err.Error(),
		})
		return
	}

	if err := model.Delete(req.Id); err != nil {
		log.Print(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code": http.StatusInternalServerError,
			"data": "数据录入错误",
			"msg":  err.Error(),
		})
		return
	}

	// 成功
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": req.Id,
		"msg":  "删除成功",
	})

}

func GetBad(ctx *gin.Context) {

}

func CountBad(ctx *gin.Context) {

}
