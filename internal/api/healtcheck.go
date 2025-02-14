package api

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	HealthCheckService interfaces.IHealthCheckService
}

func (api *HealthCheck) HealthcheckHandlerHTTP(c *gin.Context) {
	msg, err := api.HealthCheckService.HealthCheckServices()
	if err != nil {
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, err.Error(), nil)
	}

	helpers.SendResponseHTTP(c, http.StatusOK, msg, nil)
}
