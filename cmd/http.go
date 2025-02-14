package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/services"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthCheckService := &services.HealthCheck{}
	healthcheckAPI := &api.HealthCheck{HealthCheckService: healthCheckService}

	r := gin.Default()

	r.GET("/health-check", healthcheckAPI.HealthcheckHandlerHTTP)

	r.Run(":" + helpers.GetEnv("PORT", "8080"))
}
