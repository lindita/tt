package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tt.com/tt/api/app"
	http1 "tt.com/tt/internal/http"
)

//总路由
func InitRoute(r *gin.Engine, c *Api) {
	r.NoRoute(func(c *gin.Context) {
		c.PureJSON(http.StatusNotFound, http1.ApiResponse{
			ErrorCode : http.StatusNotFound,
			Msg : "not found",
		})
	})
	r.GET("/", func(c *gin.Context) {
		//输出json结果给调用方
		c.PureJSON(http.StatusOK, http1.ApiResponse{
			Msg : "hi!",
		})
	})
	app.InitRoute(r.Group("/app"), c.App)
}