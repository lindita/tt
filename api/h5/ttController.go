package h5

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
	http1.NewResponse().Success("h5 api!").Return(c)
}