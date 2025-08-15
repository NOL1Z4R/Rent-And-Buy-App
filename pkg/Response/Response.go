package Response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JSON(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

func OK(c *gin.Context, data interface{}) {
	JSON(c, http.StatusOK, "success", data)
}

func Error(c *gin.Context, data interface{}) {
	JSON(c, http.StatusBadRequest, "An error occured during process", data)
}
