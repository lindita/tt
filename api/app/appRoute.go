package app

import "github.com/gin-gonic/gin"

//app路由 可以拆分放到各目录下
func InitRoute(r *gin.RouterGroup, c *App) {

	v1 := c.V1
	v1Route := r.Group("/v1")
	v1Route.GET("tt", v1.Tt.Index)
	v1Route.GET("tt.testRedis", v1.Tt.TestRedis)
	v1Route.GET("tt.testMysql", v1.Tt.TestMysql)

	h5 := c.H5
	h5Route := r.Group("/h5")
	h5Route.GET("tt", h5.Tt.Index)
}