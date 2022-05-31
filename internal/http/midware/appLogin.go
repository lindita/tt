package midware

import (
	"github.com/gin-gonic/gin"
	"tt.com/tt/internal/constant"
	http1 "tt.com/tt/internal/http"
)

//打印耗时的中间件

func AppLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		http1.NewResponse().ErrorWithCode(constant.CodeUnLogin, "not login").Return(c)
		c.Abort()
	}
}