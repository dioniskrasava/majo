package main

import "fyne.io/fyne/v2"

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
