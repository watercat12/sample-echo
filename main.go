package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	// Echo instance
	e := echo.New()

	// Routes
	e.GET("/ok", func(context echo.Context) error {
		return context.JSON(http.StatusOK,"ok")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
