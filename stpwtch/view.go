package stopwatch

import (
	"image/color"

	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// инициализация НАЧАЛЬНОГО СОСТОЯНИЯ секундомера
func widgets_begin() {

	bgColor := theme.BackgroundColor()
	rgbaColor := color.RGBAModel.Convert(bgColor).(color.RGBA)

	var input *canvas.Text

	// ЕСЛИ ЦВЕТ ФОНА БЕЛЫЙ
	if rgbaColor == (color.RGBA{255, 255, 255, 255}) {
		input = canvas.NewText("00:00:00:000", color.RGBA{0, 0, 0, 222})
	} else {
		input = canvas.NewText("00:00:00:000", color.RGBA{255, 255, 255, 255})
	}

	input.TextSize = 35 // Увеличение размера шрифта

	btn_str := widget.NewButton("       Start       ", nil)
	btn_pause := widget.NewButton("      Pause      ", nil)
	btn_clear := widget.NewButton("      Clear      ", nil)

	wdgts.btn_str = btn_str
	wdgts.btn_pause = btn_pause
	wdgts.btn_clear = btn_clear
	wdgts.input = input

	// изначально кнопки паузы и чистки - неактивны
	btn_pause.Disable()
	btn_clear.Disable()

	btn_str.OnTapped = logic_timer("start", &flags, wdgts, &time_val)
	btn_pause.OnTapped = logic_timer("pause", &flags, wdgts, &time_val)
	btn_clear.OnTapped = logic_timer("clear", &flags, wdgts, &time_val)

	wdgts.btn_cont = container.NewHBox(btn_str, btn_pause, btn_clear)

	wdgts.cont_begin = container.NewBorder(
		nil,
		container.NewVBox(container.NewCenter(input), container.NewCenter(wdgts.btn_cont)),
		nil,
		nil,
		nil,
	)
}

func init_cont_add_act() {
	// КОНТЕНТ ДОБАВЛЕНИЯ АКТИВНОСТИ
	wdgts.input_act = widget.NewEntry()
	wdgts.input_act.SetPlaceHolder("Введите название активности...")
	wdgts.input_act_descr = widget.NewEntry()
	wdgts.input_act_descr.SetPlaceHolder("Описание активности...")
	//adding_an_activity_pad := func() { return adding_an_activity() }
	btn_add_act := widget.NewButton("Добавить активность", adding_an_activity) // ? добавлять ли в общую структуру виджетов?
	wdgts.cont_add_act = container.NewVBox(wdgts.input_act, wdgts.input_act_descr, btn_add_act)

	//adding_an_activity()
}
