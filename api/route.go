package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tt.com/tt/api/app"
	"tt.com/tt/api/h5"
	http1 "tt.com/tt/internal/http"
)

//总路由
func InitRoute(r *gin.Engine, c *Api) {
	r.NoRoute(func(c *gin.Context) {
		http1.NewResponse().HttpCode(http.StatusNotFound).ErrorWithCode( http.StatusNotFound, "not found!").Return(c)
	})
	r.GET("/", func(c *gin.Context) {
		http1.NewResponse().Success("hi!").Return(c)
	})
	app.InitRoute(r.Group("/app"), c.App)
	h5.InitRoute(r.Group("/h5"), c.H5)
}