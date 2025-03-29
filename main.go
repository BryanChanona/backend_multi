package main

import (
	temperatureDependencies "github.com/BryanChanona/backend_multi/src/Temperature/infrastructure/dependencies"
	"github.com/BryanChanona/backend_multi/src/helpers"
	"github.com/gin-gonic/gin"
	routesTemperature "github.com/BryanChanona/backend_multi/src/Temperature/infrastructure/routes"
	routesOxygen "github.com/BryanChanona/backend_multi/src/Oxygen/infrastructure/routes"
	oxygenDependencies "github.com/BryanChanona/backend_multi/src/Oxygen/infrastructure/dependencies"
	heartRateDependencies "github.com/BryanChanona/backend_multi/src/HeartRate/infrastructure/dependencies"
	routesHeartRate "github.com/BryanChanona/backend_multi/src/HeartRate/infrastructure/routes"
)

func main() {
	temperatureDependencies.Init()
	oxygenDependencies.Init()
	heartRateDependencies.Init()

	 
	r:= gin.Default()
	helpers.InitCORS(r)
	routesTemperature.Routes(r)
	routesOxygen.Routes(r)
	routesHeartRate.Routes(r)
	
	
	
   
	r.Run(":8081")

}