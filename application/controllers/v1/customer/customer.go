package application_controller_v1_customer

import (
	infra_database "gotizen/infrastructure/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)



func CustomerConroller(g *gin.RouterGroup){
	customer:=g.Group("/customer")
		{

			customer.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
				"message": "hello customer",
				})
				//  var queryCustomer,e = infra_database.GetCustomer(5345)
				// if(e){
				// 	return c.JSON(http.StatusOK, gin.H{
				// 	"404": "not customer",
				// 	})
				// }
				// return c.JSON(http.StatusOK, gin.H{
				// "data": queryCustomer.Email,
				// })
			})
			customer.GET("/:id", func(c *gin.Context) {
				var queryCustomer,e = infra_database.GetCustomer(5345)
				if(e!=nil){
					c.JSON(http.StatusNotFound, gin.H{
						"message": "oh not",
						})
				}
				userID := c.Param("id")

				if(userID ==  strconv.Itoa(queryCustomer.ID)){
					c.JSON(http.StatusOK, gin.H{
					"message": queryCustomer.Email,
					})
				}else{
					c.JSON(http.StatusNotFound, gin.H{
					"message": userID,
					})
				}
				
			})
		}
}