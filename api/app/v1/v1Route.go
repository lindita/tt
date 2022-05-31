package v1

import (
	"github.com/gin-gonic/gin"
	"tt.com/tt/internal/http/midware"
)

func InitRoute(r *gin.RouterGroup, c *V1) {

	v1Route := r.Group("/v1")
	v1LoginRoute := r.Group("/v1").Use(midware.AppLogin())

	v1Route.GET("tt", c.Tt.index)
	v1Route.GET("tt.testPanic", c.Tt.testPanic)
	v1Route.GET("tt.testRedis", c.Tt.testRedis)
	v1Route.GET("tt.testMysql", c.Tt.testMysql)
	v1LoginRoute.GET("tt.testLogin", c.Tt.testLogin)
}