package web

import (
	"github.com/labstack/echo/v4"
	"github.com/madlemon/estim8/web/template"
	"log"
	"net/http"
)

const UsernameToShort = "Your name should have at least two characters."

func SetUsername(c echo.Context) error {
	requestData := new(SignInDTO)
	if err := c.Bind(requestData); err != nil {
		log.Fatal(err)
		return err
	}

	if len(requestData.Username) < 2 {
		errorMsg := UsernameToShort
		return Render(template.UsernameForm(requestData.Username, &errorMsg), c)
	}

	authCookie := new(http.Cookie)
	authCookie.Name = "username"
	authCookie.Value = requestData.Username
	c.SetCookie(authCookie)

	return Redirect(c, requestData.RedirectTarget, http.StatusOK)
}

type SignInDTO struct {
	Username       string `form:"username"`
	RedirectTarget string `form:"redirectTarget"`
}

func ClearUsername(c echo.Context) error {
	authCookie := new(http.Cookie)
	authCookie.Name = "username"
	authCookie.Value = ""
	c.SetCookie(authCookie)

	return Redirect(c, "/", http.StatusOK)
}
