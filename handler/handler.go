package handler

import (
	"example-project/errorHandler"
	"example-project/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

const idParam = "id"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ServiceInterface
type ServiceInterface interface {
	CreateEmployees(employees []model.Employee) error
	GetEmployeeById(id string) (*model.Employee, error)
}

const (
	contentType      = "Content-Type"
	contentTypeValue = "application/json"
)

type Handler struct {
	ServiceInterface ServiceInterface
}

func NewHandler(serviceInterface ServiceInterface) *Handler {
	return &Handler{
		ServiceInterface: serviceInterface,
	}
}

func (h *Handler) SetContentType(c *gin.Context) {
	c.Header(contentType, contentTypeValue)
}

func (h *Handler) CreateEmployeeHandler(c *gin.Context) {
	var body model.CreateEmployeePayload
	err := c.BindJSON(&body)
	if err != nil {
		errorHandler.HandleError(c, errorHandler.InternalErrorInvalidPayload)
		return
	}
	err = h.ServiceInterface.CreateEmployees(body.Employees)
	if err != nil {
		errorHandler.HandleError(c, err.Error())
		return
	}
	c.JSON(http.StatusCreated, body)
}

func (h *Handler) GetEmployeeHandler(c *gin.Context) {
	id, ok := c.Params.Get(idParam)
	if !ok || id == "" {
		errorHandler.HandleError(c, errorHandler.InternalErrorEmployeeIdNotGiven)
		return
	}
	response, err := h.ServiceInterface.GetEmployeeById(id)
	if err != nil {
		errorHandler.HandleError(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, response)
}
