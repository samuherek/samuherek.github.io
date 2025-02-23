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

func error_404(c echo.Context) error {
	return renderView(c, pages.Error404Wrapper("Error", pages.Error404()))
}

func main() {
	e := echo.New()
	e.Static("/static", "static")

	e.GET("/", index)
	e.GET("/404", error_404)
	e.Logger.Fatal(e.Start(":1323"))
}
