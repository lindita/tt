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

func (c *TtController) Index(g *gin.Context) {
	str := c.service.Tt.GetTt(g)
	http1.NewResponse().Success(str).Return(g)
}

func (c *TtController) TestLogin(g *gin.Context) {
	str := c.service.Tt.GetTt(g)
	http1.NewResponse().Success(str).Return(g)
}

func (c *TtController) TestRedis(g *gin.Context) {
	data := c.service.Tt.TestRedis(g)
	http1.NewResponse().Data(data).Return(g)
}

func (c *TtController) TestMysql(g *gin.Context) {
	data := c.service.Tt.TestMysql(g)
	http1.NewResponse().Data(data).Return(g)
}

func (c *TtController) TestPanic(g *gin.Context) {
	c.service.Tt.TestPanic(g)
	http1.NewResponse().Success("success").Return(g)
}