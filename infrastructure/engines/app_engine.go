package infrastructure_engins

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MyEngine struct{
	*gin.Engine
	AppName string
}

func (engine *MyEngine) InitRouter(routerVersion string) *gin.RouterGroup{
	// if not router
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound,gin.H{"message": "Page not found"})
	})
	return engine.Group(routerVersion)
}

func(engine *MyEngine) SetVersion(version string){
	versionMiddleware := func(c *gin.Context) {
		c.Set("version", version)
	}
	engine.Use(versionMiddleware)
}