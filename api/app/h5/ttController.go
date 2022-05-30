package h5

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
	g.JSON(200, gin.H{
		"message": "h5 api!",
	})
}