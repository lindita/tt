package h5

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
	http1.NewResponse().Success("h5 api!").Return(c)
}