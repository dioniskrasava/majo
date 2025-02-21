package stopwatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func init_menu_stpwtc() *fyne.MainMenu {
	//------------------------------------------------------------------
	file_item1 := fyne.NewMenuItem("Сохранить", func() {})
	file_item2 := fyne.NewMenuItem("Открыть", func() {})

	file_menu := fyne.NewMenu("Файл", file_item1, file_item2)
	//------------------------------------------------------------------
	view_item1 := fyne.NewMenuItem("Секундомер", func() {
		wdgts.w.SetContent(wdgts.cont_begin)                   // восстановление первоначального вида секундомера
		wdgts.w.Resize(fyne.NewSize(inf_app.w_a, inf_app.h_a)) // восстановление размера окна    !!!! ПОЧЕМУ-ТО НЕ СРАБАТЫВАЕТ КОРРЕКТНО
	})
	view_item2 := fyne.NewMenuItem("Показать активность", func() {
		begin_work_db()
		wdgts.w.Resize(fyne.NewSize(inf_app.w_a, inf_app.h_a+125)) // Увеличение размера окна
		//wdgts.ent_activity.Show()
		wdgts.w.SetContent(container.NewVBox(
			container.NewCenter(container.NewGridWrap(fyne.NewSize(inf_app.w_a, 40), wdgts.ent_activity)),
			container.NewCenter(wdgts.input),
			container.NewCenter(wdgts.btn_cont)))
	})

	view_menu := fyne.NewMenu("Вид", view_item1, view_item2)
	//------------------------------------------------------------------

	edit_item1 := fyne.NewMenuItem("Добавить активность", func() {
		init_cont_add_act()
		wdgts.w.SetContent(wdgts.cont_add_act)
	})

	edit_menu := fyne.NewMenu("Редактирование", edit_item1)
	//------------------------------------------------------------------

	main_menu := fyne.NewMainMenu(file_menu, view_menu, edit_menu)

	return main_menu

}
