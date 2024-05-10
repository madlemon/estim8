package web

import (
	"github.com/labstack/echo/v4"
)

func StartWebapp(publicUrl string, port string) {
	Config = Configuration{
		PublicUrl: publicUrl,
		Port:      port,
	}

	e := echo.New()

	e.Static("/static", "web/assets")

	e.POST("/users", SetUsername)
	e.DELETE("/users", ClearUsername)
	e.GET("/", GetIndex)
	e.GET("/rooms/new", CreateRoom)
	e.GET("/rooms/:roomId", GetRoom)
	e.GET("/rooms/:roomId/estimates", GetEstimates)
	e.PUT("/rooms/:roomId/estimates", SetEstimate)
	e.DELETE("/rooms/:roomId/estimates", DeleteEstimates)
	e.PUT("/rooms/:roomId/result-visibility", SetResultVisibility)

	e.Logger.Fatal(e.Start(":" + port))
}

type Configuration struct {
	Port      string
	PublicUrl string
}

var Config Configuration
