package ui

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type UIService struct {
	server   *echo.Echo
	endpoint string
}

func New() (*UIService, error) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "static")

	e.GET("/data", func(c echo.Context) error {
		return c.String(http.StatusOK, "This content was loaded asynchronously using HTMX!")
	})

	return &UIService{
		server: e,
	}, nil
}

func (ui *UIService) Run() error {
	return ui.server.Start(ui.endpoint)
}
