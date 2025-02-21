package stopwatch

import (
	"fyne.io/fyne/v2"
)

// Приложение секундомера
func NewApp(a fyne.App) {
	/*Тут будет находится всё касаемо окна приложения*/

	inf_app.w_a = 350 // ширина
	inf_app.h_a = 120 // высота

	wdgts.w = a.NewWindow("Stopwatch")
	wdgts.w.Resize(fyne.NewSize(inf_app.w_a, inf_app.h_a)) // Увеличение размера окна
	wdgts.w.CenterOnScreen()

	widgets_begin()

	wdgts.w.SetContent(wdgts.cont_begin)
	wdgts.w.SetMainMenu(init_menu_stpwtc()) // и передача виджетов ф-ии по инициал. меню
	wdgts.w.Show()

}


