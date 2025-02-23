package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	createMainAppAndWindow()
}

func createMainAppAndWindow() {
	mainApp.a = app.NewWithID("main.journal")
	mainApp.a.Settings().SetTheme(theme.LightTheme())
	W = mainApp.a.NewWindow("Main Journal version 0.0.1")
	W.Resize(fyne.NewSize(500, 250)) // размер окна
	W.SetFixedSize(true)
	W.CenterOnScreen()

	W.SetContent(welcomeTextPrint())
	InitMainMenu()
	W.ShowAndRun()
}

func welcomeTextPrint() *fyne.Container {

	var welcomeTextLabel *widget.RichText

	bgColor := theme.BackgroundColor()
	rgbaColor := color.RGBAModel.Convert(bgColor).(color.RGBA)

	if rgbaColor == (color.RGBA{255, 255, 255, 255}) {
		welcomeTextLabel = richWelcomeText()
	} else {
		welcomeTextLabel = richWelcomeText()
	}

	contentMainWindow := container.NewHBox(welcomeTextLabel)

	return contentMainWindow
}
