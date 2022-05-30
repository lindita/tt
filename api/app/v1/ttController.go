package v1

import (
	"github.com/gin-gonic/gin"
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
	g.JSON(200, gin.H{
		"message": str,
	})
}