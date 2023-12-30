package main

import (
	infrastructure_engins "gotizen/infrastructure/engines"

	application_controller_v1_customer "gotizen/application/controllers/v1/customer"
	application_controller_v1_user "gotizen/application/controllers/v1/user"

	"github.com/gin-gonic/gin"
)

func main() {
  // setup config
 	appEngine  := &infrastructure_engins.MyEngine{AppName: "Main App",Engine: gin.Default()}
	appEngine.SetVersion("0.0.1")
  gin.SetMode(gin.DebugMode)
  routerGroup := appEngine.InitRouter("/v1")

  // push router
  application_controller_v1_customer.CustomerConroller(routerGroup)
  application_controller_v1_user.UserConroller(routerGroup)

  
  // just running
  appEngine.Run("127.0.0.1:3000")
}