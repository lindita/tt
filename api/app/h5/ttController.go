package h5

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
	g.PureJSON(http.StatusOK, http1.ApiResponse{
		Msg: "h5 api!",
	})
}