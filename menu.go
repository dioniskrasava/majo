package main

import (
	"fmt"
	"majo/fixact"
	stopwatch "majo/stpwtch"

	"fyne.io/fyne/v2"
)

type Lang_data struct {
	utils_menu_name       string
	utils_menu_name_i1    string // название первого приложения
	utils_menu_name_i2    string // название второго приложения
	sett_menu_name        string // меню настроек
	sett_item1_name       string
	sett_item2_name       string
	sett_item1_lang1_name string // первый язык
	sett_item1_lang2_name string
}

var eng_l Lang_data = Lang_data{
	utils_menu_name:       "Utils",
	utils_menu_name_i1:    "Stopwatch",
	utils_menu_name_i2:    "FixAct",
	sett_menu_name:        "Settings",
	sett_item1_name:       "Language",
	sett_item2_name:       "Theme",
	sett_item1_lang1_name: "Russian",
	sett_item1_lang2_name: "English",
}

var rus_l Lang_data = Lang_data{
	utils_menu_name:       "Инструменты",
	utils_menu_name_i1:    "Секундомер",
	utils_menu_name_i2:    "Активности",
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

func InitMainMenu() {
	utils_menu, sett_menu := creatingMenuItems()
	main_menu = fyne.NewMainMenu(utils_menu, sett_menu)
	W.SetMainMenu(main_menu)
}

// Function to update the menu when language changes
func updateMenu() {
	if main_menu == nil || W == nil {
		return
	}
	utils_menu, sett_menu := creatingMenuItems()
	main_menu.Items[0] = utils_menu
	main_menu.Items[1] = sett_menu
	main_menu.Refresh()
}

func creatingMenuItems() (*fyne.Menu, *fyne.Menu) {

	utils_stopwatch := fyne.NewMenuItem(set_lang.utils_menu_name_i1, func() { stopwatch.NewApp(mainApp.a) })
	utils_fixact := fyne.NewMenuItem(set_lang.utils_menu_name_i2, func() { fixact.NewApp(mainApp.a) })
	utils_separator := fyne.NewMenuItemSeparator()

	// UTILS main-item
	utils_menu := fyne.NewMenu(set_lang.utils_menu_name, utils_stopwatch, utils_fixact, utils_separator)
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

	// SETTINGS main-item
	sett_menu := fyne.NewMenu(set_lang.sett_menu_name, sett_item1, sett_item2)

	return utils_menu, sett_menu
}
