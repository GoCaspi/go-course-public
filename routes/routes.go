package routes

import (
	"github.com/gin-gonic/gin"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . HandlerInterface
type HandlerInterface interface {
	CreateEmployeeHandler(c *gin.Context)
	GetEmployeeHandler(c *gin.Context)
	SetContentType(c *gin.Context)
}

var Handler HandlerInterface

func CreateRoutes(group *gin.RouterGroup) {
	employees := group.Group("/employees")
	employees.Use(Handler.SetContentType)
	employees.GET("/:id", Handler.GetEmployeeHandler)
	employees.POST("/", Handler.CreateEmployeeHandler)
}
