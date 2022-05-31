package midware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"tt.com/tt/internal/exception"
	http1 "tt.com/tt/internal/http"
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
				c.JSON(http.StatusOK, http1.ApiResponse{
					ErrorCode: exception.ErrorCode,
					Msg:       exception.Msg,
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
