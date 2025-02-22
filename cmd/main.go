package main

import (
	"samuherek/pages"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func index(c echo.Context) error {
	return renderView(c, pages.HomeWrapper("Home", pages.Home()))
}

func main() {
	e := echo.New()
	e.Static("/static", "static")

	e.GET("/", index)
	e.Logger.Fatal(e.Start(":1323"))
}
