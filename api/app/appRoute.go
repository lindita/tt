package app

import (
	"github.com/gin-gonic/gin"
	v1 "tt.com/tt/api/app/v1"
)

//app路由 可以拆分放到各目录下
func InitRoute(r *gin.Engine, c *App) {
	v1.InitRoute(r.Group("/app"), c.V1)
}