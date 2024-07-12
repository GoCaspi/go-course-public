package errorHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, message string) {
	c.AbortWithStatusJSON(toHttpStatus(message), gin.H{
		"errorMessage": toExternalErrorMessage(message),
	})
}

func toHttpStatus(message string) int {
	switch message {
	case InternalErrorNoEmployeeFound:
		return http.StatusNotFound
	case InternalErrorInvalidPayload:
		return http.StatusBadRequest
	case InternalErrorEmployeeIdNotGiven, InternalErrorEmployeeIdNotUnique:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func toExternalErrorMessage(message string) string {
	switch message {
	case InternalErrorNoEmployeeFound:
		return ExternalErrorEmployeeNotFound
	case InternalErrorInvalidPayload:
		return ExternalErrorInvalidPayload
	case InternalErrorEmployeeIdNotGiven:
		return ExternalErrorIdMissingInURL
	case InternalErrorEmployeeIdNotUnique:
		return ExternalErrorEmployeeIdNotUnique
	default:
		return ExternalErrorUnknown
	}
}
