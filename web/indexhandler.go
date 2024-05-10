package web

import (
	"github.com/labstack/echo/v4"
	"github.com/madlemon/estim8/web/template"
)

func GetIndex(c echo.Context) error {
	authCookie, err := c.Cookie("username")
	if CookieHasNoValue(err, authCookie) {
		return Render(template.LoginView(), c)
	}

	component := template.LandingView()
	return Render(component, c)
}
