package http

import (
	"sober_driver/internal/config"
	"sober_driver/internal/service"
	v1 "sober_driver/internal/transport/http/v1"
	"sober_driver/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Service    *service.Service
	configSwag *config.SwaggerConfig
}

func NewHandler(service *service.Service, configSwag *config.SwaggerConfig) *Handler {
	return &Handler{
		Service:    service,
		configSwag: configSwag,
	}
}

func (hr *Handler) Init() *gin.Engine {
	g := utils.InitGin()
	//iniApi(g.Router())
	//return g.Engine()
	v1.InitDocumentationServer(hr.configSwag, g.Engine())
	hr.initApi(g.Router())
	return g.Engine()
}

func (hr *Handler) initApi(router *gin.RouterGroup) {
	v1Handler := v1.NewHandler(hr.Service)
	v1Handler.Init(router.Group("/v1"))
}

// //
