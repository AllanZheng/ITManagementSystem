package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
	api "system/api"
	model "system/model"
)

func main() {
	r := gin.Default() //初始化gin默认框架
	model.Init()
	//defer db.Close()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/fee/create", api.CreateRecord)
	r.POST("/fee/update", api.UpdateRecord)
	r.POST("/fee/delete", api.DeleteRecord)
	r.POST("/fee/get", api.GetRecord)
	r.POST("/bad/create", api.CreateBad)
	r.POST("/bad/update", api.UpdateBad)
	r.POST("/bad/delete", api.DeleteBad)
	r.POST("/bad/get", api.GetBad)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
