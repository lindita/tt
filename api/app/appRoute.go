package app

import (
	"github.com/gin-gonic/gin"
	midware2 "tt.com/tt/internal/http/midware"
)

//app路由 可以拆分放到各目录下
func InitRoute(r *gin.RouterGroup, c *App) {

	v1 := c.V1
	v1Route := r.Group("/v1")
	v1LoginRoute := r.Group("/v1").Use(midware2.AppLogin())
	v1Route.GET("tt", v1.Tt.Index)
	v1Route.GET("tt.testPanic", v1.Tt.TestPanic)
	v1Route.GET("tt.testRedis", v1.Tt.TestRedis)
	v1Route.GET("tt.testMysql", v1.Tt.TestMysql)
	v1LoginRoute.GET("tt.testLogin", v1.Tt.TestLogin)

	h5 := c.H5
	h5Route := r.Group("/h5")
	h5Route.GET("tt", h5.Tt.Index)
}