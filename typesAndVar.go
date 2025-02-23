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

func richWelcomeText() *widget.RichText {
	markdownText := "# **Main Journal**\n\n" +
		"	" +
		"приложение для контроля активностей/деятельностей\n" +
		"---\n" +
		"В дальнейшем планируется добавить приложение для контроля : \n\n" +
		"**за питанием, финансами, ежедневными рутинами.**\n" +
		" \n" +
		" \n"

	// Создание RichText из Markdown
	richText := widget.NewRichTextFromMarkdown(markdownText)
	return richText
}
