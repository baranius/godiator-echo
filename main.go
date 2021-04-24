package main

import (
	"github.com/baranius/godiator"
	"github.com/baranius/godiator-echo/resources"
	"github.com/baranius/godiator-echo/resources/healthcheck/healthcheck_handlers"
	"github.com/baranius/godiator-echo/utils/pipelines"
	"github.com/labstack/echo"
)

func main() {
	//Create Echo
	e := echo.New()

	//Get Godiator
	g := godiator.GetInstance()

	//RegisterPipeline
	g.RegisterPipeline(&pipelines.ValidationPipeline{})

	//Register Handlers
	g.Register(&healthcheck_handlers.HealthCheckRequest{}, healthcheck_handlers.NewHealthCheckHandler)

	//Register Resources
	resources.Register(e)

	//Register Router
	e.Logger.Fatal(e.Start(":8080"))
}
