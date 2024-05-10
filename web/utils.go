package web

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Render(component templ.Component, c echo.Context) error {
	return component.Render(c.Request().Context(), c.Response())
}

func Redirect(c echo.Context, url string, code int) error {
	c.Response().Header().Add("HX-Redirect", url)
	return c.NoContent(code)
}

func CookieHasNoValue(err error, authCookie *http.Cookie) bool {
	return err != nil || authCookie == nil || authCookie.Value == ""
}
