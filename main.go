package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var W fyne.Window

type globalObjectsApplication struct {
	a fyne.App
}

var mainApp globalObjectsApplication

func main() {

	mainApp.a = app.NewWithID("main.journal")
	mainApp.a.Settings().SetTheme(theme.LightTheme())
	W := mainApp.a.NewWindow("Main Journal version 0.0.1")
	W.Resize(fyne.NewSize(500, 250)) // размер окна
	W.SetFixedSize(true)
	W.CenterOnScreen()

	welcomeText := `Main Journal - приложение для контроля активностей/деятельностей.
	В дальнейшем планируется добавить приложение для контроля
за питанием, финансами, ежедневными рутинами.

	Под вопросом:
1) Должны ли все микроприложения быть отдельными приложениями (окнами) или являться
частью одного окна? (Наверное комбинированный вариант - какие-то утилиты
отдельными)`

	welcomeTextLabel := widget.NewLabel(welcomeText)

	contentMainWindow := container.NewHBox(welcomeTextLabel)

	W.SetContent(contentMainWindow)
	W.SetMainMenu(InitMenu(W))
	W.ShowAndRun()

}
