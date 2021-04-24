package healthcheck

import (
	"github.com/baranius/godiator"
	"github.com/baranius/godiator-echo/resources/healthcheck/healthcheck_handlers"
	"github.com/labstack/echo"
	"net/http"
)

type Resource struct {
	gdtr godiator.IGodiator
}

func NewResource() Resource {
	return Resource{gdtr: godiator.GetInstance()}
}

func (h *Resource) GetHealthCheck(c echo.Context) error {
	req := &healthcheck_handlers.HealthCheckRequest{}
	response, err := h.gdtr.Send(req, c)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, response)
}