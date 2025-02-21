package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var W fyne.Window

type globalObjectsApplication struct {
	a fyne.App
}

var mainApp globalObjectsApplication

var welcomeText string = `Main Journal - приложение для контроля активностей/деятельностей.

	В дальнейшем планируется добавить приложение для контроля

за питанием, финансами, ежедневными рутинами.

	Под вопросом:

1) Должны ли все микроприложения быть отдельными приложениями (окнами) или являться
частью одного окна? (Наверное комбинированный вариант - какие-то утилиты
отдельными)`

/*
var welcomeText string = "Main Journal - приложение для контроля активностей/деятельностей.\n" +
	"В дальнейшем планируется добавить приложение для контроля\n" +
	"за питанием, финансами, ежедневными рутинами.\n" +

	"Под вопросом:\n" +
	"1) Должны ли все микроприложения быть отдельными приложениями (окнами) или являться\n" +
	"частью одного окна? (Наверное комбинированный вариант - какие-то утилиты\n" +
	"отдельными)\n"
*/

func richWelcomeText() *widget.RichText {
	// Многострочный текст с форматированием
	richText := widget.NewRichText(
		&widget.TextSegment{Text: "Это пример многострочного текста.\n"},
		&widget.TextSegment{Text: "Он поддерживает ", Style: widget.RichTextStyleInline},
		&widget.TextSegment{Text: "разное форматирование", Style: widget.RichTextStyleBold},
		&widget.TextSegment{Text: " и перенос строк.", Style: widget.RichTextStyleInline},
	)
	return richText
}
