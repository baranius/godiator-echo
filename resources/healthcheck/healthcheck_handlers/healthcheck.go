package healthcheck_handlers

import "github.com/labstack/echo"

type (
	HealthCheckRequest struct {
	}

	HealthCheckResponse struct {
		RequestMethod string `json:"request-method"`
		Success       bool   `json:"success"`
	}

	HealthCheckHandler struct {
	}
)

func NewHealthCheckHandler() interface{} {
	return &HealthCheckHandler{}
}

func (h *HealthCheckHandler) Handle(request *HealthCheckRequest, c echo.Context) (*HealthCheckResponse, error) {
	return &HealthCheckResponse{Success: true, RequestMethod: c.Request().Method}, nil
}
