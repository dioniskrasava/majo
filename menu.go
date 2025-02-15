package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
)

type Lang_data struct {
	app_menu_name         string
	app_menu_name_i1      string
	sett_menu_name        string
	sett_item1_name       string
	sett_item2_name       string
	sett_item1_lang1_name string
	sett_item1_lang2_name string
}

var eng_l Lang_data = Lang_data{
	app_menu_name:         "Applications",
	app_menu_name_i1:      "Stopwatch",
	sett_menu_name:        "Settings",
	sett_item1_name:       "Language",
	sett_item2_name:       "Theme",
	sett_item1_lang1_name: "Russian",
	sett_item1_lang2_name: "English",
}

var rus_l Lang_data = Lang_data{
	app_menu_name:         "Приложения",
	app_menu_name_i1:      "Секундомер",
	sett_menu_name:        "Настройки",
	sett_item1_name:       "Язык",
	sett_item2_name:       "Тема",
	sett_item1_lang1_name: "Русский",
	sett_item1_lang2_name: "Английский",
}

// Current selected language
var set_lang Lang_data = eng_l

// Main menu variable declared globally to be accessible by all functions
var main_menu *fyne.MainMenu

func setRusLang() {
	fmt.Println("rus")
	set_lang = rus_l
	updateMenu()
}

func setEngLang() {
	fmt.Println("eng")
	set_lang = eng_l
	updateMenu()
}

func InitMenu(window fyne.Window) *fyne.MainMenu {
	W = window
	app_menu, sett_menu := creatingMenuItems()
	main_menu = fyne.NewMainMenu(app_menu, sett_menu)
	window.SetMainMenu(main_menu)

	return main_menu
}

// Function to update the menu when language changes
func updateMenu() {
	if main_menu == nil || W == nil {
		return
	}
	app_menu, sett_menu := creatingMenuItems()
	main_menu.Items[0] = app_menu
	main_menu.Items[1] = sett_menu
	main_menu.Refresh()
}

func creatingMenuItems() (*fyne.Menu, *fyne.Menu) {
	// --stopwatch (1-1)
	app_item1 := fyne.NewMenuItem(set_lang.app_menu_name_i1, func() {})
	app_separator := fyne.NewMenuItemSeparator()
	app_item_quit := fyne.NewMenuItem("QuВыйти из программыit", func() { os.Exit(0) })
	// application (1)
	app_menu := fyne.NewMenu(set_lang.app_menu_name, app_item1, app_separator, app_item_quit)
	//-----------------------------------------------------------------------------
	// --language (2-1)
	sett_item1 := fyne.NewMenuItem(set_lang.sett_item1_name, func() {})
	// -- --rus (2-1-1)
	sett_item1_lang1 := fyne.NewMenuItem(set_lang.sett_item1_lang1_name, setRusLang)
	// -- --eng (2-1-2)
	sett_item1_lang2 := fyne.NewMenuItem(set_lang.sett_item1_lang2_name, setEngLang)
	//
	sett_item1.ChildMenu = fyne.NewMenu("*-*-*",
		sett_item1_lang1,
		sett_item1_lang2)

	// --theme (2-2)
	sett_item2 := fyne.NewMenuItem(set_lang.sett_item2_name, func() {})

	// settings (2)
	sett_menu := fyne.NewMenu(set_lang.sett_menu_name, sett_item1, sett_item2)

	return app_menu, sett_menu
}
