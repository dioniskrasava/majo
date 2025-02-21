package stopwatch

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
)

// СТРУКТУРА ОПИСЫВАЮЩАЯ ФЛАГИ ПРИЛОЖЕНИЯ СЕКУНДОМЕРА
type flags_app struct {
	start_pressed bool
	paused        bool
	begin_flag    bool
}

var flags flags_app = flags_app{false, false, true}

// структура кнопок
type widgets_app struct {
	btn_str         *widget.Button
	btn_pause       *widget.Button
	btn_clear       *widget.Button
	input           *canvas.Text
	cont_begin      *fyne.Container
	w               fyne.Window
	btn_cont        *fyne.Container
	ent_activity    *widget.Select
	cont_add_act    *fyne.Container
	input_act       *widget.Entry
	input_act_descr *widget.Entry
}

var wdgts widgets_app

// структура кол-ва времени в таймере
type time_values struct {
	h           int
	m           int
	s           int
	ms          int
	msTimer     int
	msStartUnix int
	timePause   int
	times       string
}

var time_val time_values = time_values{
	times: "00:00:00:000",
}

type information_app struct {
	w_a float32
	h_a float32
	DB  *gorm.DB
}

var inf_app information_app

type Activity struct {
	Id            uint
	Name_activity string
	Description   string
}
