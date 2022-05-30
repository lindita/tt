package app

import "github.com/gin-gonic/gin"

//app路由
func InitRoute(r *gin.RouterGroup, c *AppV1) {
	loginRouter := r.Group("/tt")
	loginRouter.GET("", c.Tt.Index)
}