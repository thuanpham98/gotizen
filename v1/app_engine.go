package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gotizen/v1/customer"
)

type MyEngine struct{
	*gin.Engine
	AppName string
}

func (engine *MyEngine) InitRouter(){
	// if not router
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound,gin.H{"message": "Page not found"})
	})
	
	r:=engine.Group("/v1")
	{

		user:=r.Group("/user")
		{
			user.GET("/detail", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
				"message": "hello user",
				})
			})
		}
		customer.CustomerConroller(r)
		
	}
}