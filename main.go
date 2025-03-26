package main

import (
	temperatureDependencies "github.com/BryanChanona/backend_multi/src/Temperature/infrastructure/dependencies"
	userDependencies "github.com/BryanChanona/backend_multi/src/User/infrastructure/dependencies"
	routesUser "github.com/BryanChanona/backend_multi/src/User/infrastructure/routes"
	"github.com/BryanChanona/backend_multi/src/helpers"
	"github.com/gin-gonic/gin"
	routesTemperature "github.com/BryanChanona/backend_multi/src/Temperature/infrastructure/routes"
)

func main() {
	userDependencies.Init()
	temperatureDependencies.Init()
	 
	r:= gin.Default()
	helpers.InitCORS(r)
	routesUser.Routes(r)
	routesTemperature.Routes(r)
	
	
	
   
	r.Run(":8081")

}