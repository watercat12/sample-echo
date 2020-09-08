package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	//// Echo instance
	//e := echo.New()
	//
	//// Routes
	//e.GET("/ok", func(context echo.Context) error {
	//	return context.JSON(http.StatusOK,"ok")
	//})
	//
	//// Start server
	//e.Logger.Fatal(e.Start(":1323"))

	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome :)")
		}),
	))

	w.ShowAndRun()
}
