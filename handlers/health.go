package handlers

import (
	"net/http"

	"github.com/evamdf/api-project/models"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "API is running",
	})
}
