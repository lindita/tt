package h5

import (
	"github.com/gin-gonic/gin"
)

//app路由 可以拆分放到各目录下
func InitRoute(r *gin.Engine, c *H5) {
	group := r.Group("/h5")
	group.GET("tt", c.Tt.index)
}