package app

import "github.com/gin-gonic/gin"

//app路由
func InitRoute(r *gin.RouterGroup, c *App) {

	v1 := c.V1
	v1Route := r.Group("/v1")
	v1Route.GET("tt", v1.Tt.Index)

	h5 := c.H5
	h5Route := r.Group("/h5")
	h5Route.GET("tt", h5.Tt.Index)
}