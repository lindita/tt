package v1

import (
	"github.com/gin-gonic/gin"
	http1 "tt.com/tt/internal/http"
	"tt.com/tt/internal/service"
)

type ttController struct {
	service *service.Service
}

func NewTtController(s *service.Service) *ttController {
	return  &ttController{
		service : s,
	}
}

func (t *ttController) index(c *gin.Context) {
	str := t.service.Tt.GetTt(c)
	http1.NewResponse().Success(str).Return(c)
}

func (t *ttController) testLogin(c *gin.Context) {
	str := t.service.Tt.GetTt(c)
	http1.NewResponse().Success(str).Return(c)
}

func (t *ttController) testRedis(c *gin.Context) {
	data := t.service.Tt.TestRedis(c)
	http1.NewResponse().Data(data).Return(c)
}

func (t *ttController) testMysql(c *gin.Context) {
	data := t.service.Tt.TestMysql(c)
	http1.NewResponse().Data(data).Return(c)
}

func (t *ttController) testPanic(c *gin.Context) {
	t.service.Tt.TestPanic(c)
	http1.NewResponse().Success("success").Return(c)
}