package main

import (
	"fyne.io/fyne/v2/app"
)

const Version = "0.0.1"

func main() {
	app := app.New()
	window := app.NewWindow("Huy")

	window.ShowAndRun()
}
