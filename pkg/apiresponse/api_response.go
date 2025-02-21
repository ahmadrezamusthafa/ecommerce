package apiresponse

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   error       `json:"error,omitempty"`
}

func Success(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func Error(c *gin.Context, statusCode int, message string, err error) {
	c.JSON(statusCode, APIResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	})
}
