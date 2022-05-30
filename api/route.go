package api

import (
	"github.com/gin-gonic/gin"
	"tt.com/tt/api/app"
	appV1 "tt.com/tt/api/app/v1"
)

//总路由
func InitRoute(r *gin.Engine, c *appV1.Controller) {
	r.GET("/", func(c *gin.Context) {
		//输出json结果给调用方
		c.JSON(200, gin.H{
			"message": "hi!",
		})
	})
	app.InitRoute(r.Group("/app/v1"), c)
}