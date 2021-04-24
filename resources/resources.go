package resources

import (
	"github.com/baranius/godiator-echo/resources/healthcheck"
	"github.com/labstack/echo"
)

func Register(e *echo.Echo) {
	v1 := e.Group("/v1")

	// Health Check
	healthCheckResource := healthcheck.NewResource()
	healthCheckGroup := v1.Group("/health-check")
	healthCheckGroup.GET("", healthCheckResource.GetHealthCheck)
}