package server

import (
	"example-project/handler"
	"example-project/routes"
	"example-project/service"
	"github.com/gin-gonic/gin"
)

func SetupEngine() *gin.Engine {
	engine := gin.Default()
	routes.CreateRoutes(&engine.RouterGroup)
	return engine
}

func SetupService(dbClient service.DatabaseInterface) {
	employeeService := service.NewEmployeeService(dbClient)
	routes.Handler = handler.NewHandler(employeeService)
}
