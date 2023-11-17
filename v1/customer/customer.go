package customer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerConroller(g *gin.RouterGroup){
	user:=g.Group("/customer")
		{
			user.GET("/detail", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
				"message": "hello user",
				})
			})
		}
}