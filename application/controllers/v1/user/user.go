package application_controller_v1_user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserConroller(g *gin.RouterGroup){
	user:=g.Group("/user")
		{
			user.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
				"message": "hello user",
				})
			})
		}
}