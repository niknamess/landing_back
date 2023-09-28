package utils

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

type Gin struct {
	engine *gin.Engine
	router *gin.RouterGroup
}

func (g *Gin) Engine() *gin.Engine {
	return g.engine
}
func (g *Gin) Router() *gin.RouterGroup {
	return g.router
}

func InitGin() *Gin {

	engine := gin.New()

	engine.Use(
		recoveryMiddleware,
		setTimeMiddleware,
		gin.Logger(),
	)
	router := engine.Group("/api")

	return &Gin{engine: engine,
		router: router}
}

func recoveryMiddleware(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	c.Next()
}

func setTimeMiddleware(c *gin.Context) {
	c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "time_request", time.Now()))
	c.Next()
}

func corsMiddleware(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")

	c.Next()
}
