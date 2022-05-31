package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	g.PureJSON(http.StatusOK, http1.ApiResponse{
		Msg: str,
	})
}

func (c *TtController) TestRedis(g *gin.Context) {
	str := c.service.Tt.TestRedis(g)
	g.PureJSON(http.StatusOK, http1.ApiResponse{
		Msg: str,
	})
}

func (c *TtController) TestMysql(g *gin.Context) {
	str := c.service.Tt.TestMysql(g)
	g.PureJSON(http.StatusOK, http1.ApiResponse{
		Msg: str,
	})
}

func (c *TtController) TestPanic(g *gin.Context) {
	str := c.service.Tt.TestPanic(g)
	g.PureJSON(http.StatusOK, http1.ApiResponse{
		Msg: str,
	})
}