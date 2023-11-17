package main

import (
	v1 "gotizen/v1"

	"github.com/gin-gonic/gin"
)

func main() { 
	// my engine v1
 appEngine  := &v1.MyEngine{AppName: "Main App",Engine: gin.Default()}
  appEngine.InitRouter()

  appEngine.Run("localhost:6969")
}