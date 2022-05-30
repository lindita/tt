package app

import "github.com/gin-gonic/gin"

//app路由
func InitRoute(r *gin.RouterGroup, c *AppV1) {
	route := r.Group("/tt")
	route.GET("", c.Tt.Index)
}