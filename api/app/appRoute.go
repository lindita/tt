package app

import (
	"github.com/gin-gonic/gin"
	appV1 "test.com/tt/api/app/v1"
)


func InitRoute(r *gin.RouterGroup, c *appV1.Controller) {
	loginRouter := r.Group("/tt")
	loginRouter.GET("/index", c.Tt.Index)

}