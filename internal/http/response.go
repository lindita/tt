package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tt.com/tt/internal/constant"
)

type ApiResponse struct {
	ErrorCode int         `json:"errorCode"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data"`
}

type Response struct {
	ApiResponse ApiResponse `json:"apiResponse"`
	StatusCode  int         `json:"statusCode"`
}

func NewResponse() *Response {
	return &Response{
		StatusCode: http.StatusOK,
	}
}

func (r *Response) HttpCode(status int) *Response {
	r.StatusCode = status
	return r
}

func (r *Response) Success(msg string) *Response {
	r.ApiResponse.Msg = msg
	return r
}

func (r *Response) Data(data interface{}) *Response {
	r.ApiResponse.Data = data
	return r
}

func (r *Response) Error(msg string) *Response {
	r.ApiResponse.ErrorCode = constant.CodeError
	r.ApiResponse.Msg = msg
	return r
}

func (r *Response) ErrorWithCode(errorCode int, msg string) *Response {
	r.ApiResponse.ErrorCode = errorCode
	r.ApiResponse.Msg = msg
	return r
}

func (r *Response) Return(c *gin.Context) {
	c.Get("http.isResponse")
	if  _, exist := c.Get("http.isResponse"); !exist {
		c.PureJSON(r.StatusCode, r.ApiResponse)
		c.Set("http.isResponse", 1)
	}
}
