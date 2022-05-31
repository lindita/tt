package v1

import (
	"github.com/gin-gonic/gin"
	http1 "tt.com/tt/internal/http"
	"tt.com/tt/internal/service"
)

type TtController struct {
	service *service.Service
}

func NewTtController(s *service.Service) *TtController {
	return  &TtController{
		service : s,
	}
}

func (t *TtController) Index(c *gin.Context) {
	str := t.service.Tt.GetTt(c)
	http1.NewResponse().Success(str).Return(c)
}

func (t *TtController) TestLogin(c *gin.Context) {
	str := t.service.Tt.GetTt(c)
	http1.NewResponse().Success(str).Return(c)
}

func (t *TtController) TestRedis(c *gin.Context) {
	data := t.service.Tt.TestRedis(c)
	http1.NewResponse().Data(data).Return(c)
}

func (t *TtController) TestMysql(c *gin.Context) {
	data := t.service.Tt.TestMysql(c)
	http1.NewResponse().Data(data).Return(c)
}

func (t *TtController) TestPanic(c *gin.Context) {
	t.service.Tt.TestPanic(c)
	http1.NewResponse().Success("success").Return(c)
}