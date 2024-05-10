package main

import (
	"github.com/madlemon/estim8/web"
	"os"
)

func main() {
	var publicUrl = os.Getenv("PUBLIC_URL")
	if publicUrl == "" {
		publicUrl = "http://localhost:1234"
	}

	var port = os.Getenv("PORT")
	if port == "" {
		port = "1234"
	}

	web.StartWebapp(publicUrl, port)
}
