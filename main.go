package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

var W fyne.Window

func main() {

	a := app.New()
	a.Settings().SetTheme(theme.LightTheme())
	W := a.NewWindow("Main Journal version 0.0.1")
	W.Resize(fyne.NewSize(500, 250)) // размер окна
	W.SetFixedSize(true)
	W.CenterOnScreen()

	W.SetMainMenu(InitMenu(W))
	W.ShowAndRun()

}
