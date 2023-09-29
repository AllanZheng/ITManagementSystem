package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type DBcontroller struct {
	DB *gorm.DB
}

func (db DBcontroller) CreateRecord(ctx *gin.Context) {
	var requestPost CreatePostRequest
	// 数据验证
	if err := ctx.ShouldBind(&requestPost); err != nil {
		log.Print(err.Error())
		ctx.JSON(500, gin.H{"msg": "create FAILED"})
		return
	}
	NewRecord := Record{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	if err := db.DB.Create(&NewRecord).Error; err != nil {
		err.Error(err)
	}
	ctx.JSON(200, gin.H{"msg": "create successful"})

}
