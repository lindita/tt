package h5

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
	g.JSON(200, gin.H{
		"message": "h5 api!",
	})
}