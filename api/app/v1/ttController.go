package v1

import (
	"github.com/gin-gonic/gin"
	"tt.com/tt/internal/services"
)

type TtController struct {
	service *services.Service
}

func NewTtController(s *services.Service) *TtController {
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