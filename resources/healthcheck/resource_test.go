package healthcheck

import (
	"github.com/baranius/godiator"
	"github.com/baranius/godiator-echo/resources/healthcheck/healthcheck_handlers"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ResourceTestSuite struct {
	suite.Suite
}

func TestResourceSuite(t *testing.T){
	suite.Run(t, new(ResourceTestSuite))
}

func (s *ResourceTestSuite) Test_HealthCheck() {
	// Given
	e := echo.New()

	httpRequest := httptest.NewRequest(http.MethodGet, "/v1/health-check", nil)
	httpRecorder := httptest.NewRecorder()

	context := e.NewContext(httpRequest, httpRecorder)

	mockGodiator := godiator.MockGodiator{}
	mockGodiator.OnSend = func(request interface{}, params ...interface{}) (interface{}, error) {
		return &healthcheck_handlers.HealthCheckResponse{
			RequestMethod: "GET",
			Success:       true,
		}, nil
	}

	resource := Resource{gdtr: &mockGodiator}

	// When
	err := resource.GetHealthCheck(context)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, httpRecorder.Code)
}
