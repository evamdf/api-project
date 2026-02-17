package routes

import (
	"github.com/evamdf/api-project/handlers"
	"github.com/evamdf/api-project/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.Use(middleware.Logger())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/health", handlers.HealthCheck)
	}
}
