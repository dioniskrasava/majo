package fixact

import (
	"fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

// ВЕРСИЯ  V 1.1.001 - ДОБАВИТЬ ФИЧУ (ВЫБОР МЕСТА СОХРАНИЕНИЯ БАЗЫ ДАННЫХ)

func NewApp(a fyne.App) {

	w := a.NewWindow("FixAct")
	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	w.SetFixedSize(true)

	// Получаем объект Preferences
	//prefs := a.Preferences()
	/*
			Preferences предоставляет методы для работы с различными типами данных:
		1. Строки

		    SetString(key string, value string): Сохраняет строку по ключу.
		    String(key string) string: Возвращает строку по ключу. Если ключ отсутствует, возвращает пустую строку.
		    StringWithFallback(key string, fallback string) string: Возвращает строку по ключу или значение fallback, если ключ отсутствует.

			6. Удаление данных

		    RemoveValue(key string): Удаляет значение по ключу.
	*/

	createDB()
	defer db.Close()

	createTableInDB()
	content := createInterfaceApp()

	// Установка контента в окно
	w.SetContent(content)
	w.SetMainMenu(createMenu(w))
	w.Show()
}
