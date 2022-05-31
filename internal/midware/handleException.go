package midware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"tt.com/tt/internal/exception"
)

//前置于handleRecovery，只拦截ErrorCode>0的panic,errCode=0会被当成当成系统异常抛出
func HandleException() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var exception exception.Exception
				jsonErr := json.Unmarshal([]byte(string(err.(string))), &exception)
				if jsonErr != nil {
					panic(err.(string))
				}
				if exception.ErrorCode == 0 {
					panic(err.(string))
				}
				c.JSON(200, gin.H{
					"errorCode": exception.ErrorCode,
					"msg": exception.Msg,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}

