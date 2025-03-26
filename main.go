package main

import (
	"github.com/BryanChanona/backend_multi/src/User/infrastructure/dependencies"
	"github.com/BryanChanona/backend_multi/src/User/infrastructure/routes"
	"github.com/BryanChanona/backend_multi/src/helpers"
	"github.com/gin-gonic/gin"
)

func main() {
	dependencies.Init()
	 
	r:= gin.Default()
	helpers.InitCORS(r)
	routes.Routes(r)
	
	
   
	r.Run(":8081")

}