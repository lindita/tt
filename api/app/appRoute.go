package app

import (
	"github.com/gin-gonic/gin"
	appV1 "tt.com/tt/api/app/v1"
)

//app路由
func InitRoute(r *gin.RouterGroup, c *appV1.Controller) {
	loginRouter := r.Group("/tt")
	loginRouter.GET("", c.Tt.Index)
}