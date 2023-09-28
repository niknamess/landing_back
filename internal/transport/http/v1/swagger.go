package v1

import (
	"sober_driver/internal/config"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// This is what I needed to add for it to work, this is "docs" in the root of my application
	// generated with swag init
	_ "sober_driver/cmd/sober-driver/docs"
)

func InitDocumentationServer(conf *config.SwaggerConfig, g *gin.Engine) {
	if !conf.Enabled {
		return
	}
	//g.StaticFS("/be/static", http.Dir(conf.DirPath))
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.URL(conf.URL)))
}
