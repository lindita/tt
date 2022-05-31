package midware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//打印耗时的中间件

func AppLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("appLogin")
	}
}