package stopwatch

import (
	"fmt"
	"time"
)

// логика работы таймера
func logic_timer(s string, f *flags_app, w widgets_app, t *time_values) func() {
	// функция обновления состояния кнопок
	updateButtonsState := func(startPressed bool, paused bool) {
		w.btn_str.Disable()
		w.btn_pause.Disable()
		w.btn_clear.Disable()

		if startPressed && !paused {
			w.btn_pause.Enable()
			w.btn_clear.Enable()
		} else if !startPressed && paused {
			w.btn_str.Enable()
			w.btn_clear.Enable()
		} else if !startPressed && !paused {
			w.btn_str.Enable()
		}
	}

	// СТАРТ
	start := func() {
		if !f.start_pressed {
			f.start_pressed = true
			w.btn_clear.Enable()

			if f.paused {
				t.msStartUnix = int(time.Now().UnixMilli()) - t.timePause
			}

			f.paused = false

			if f.begin_flag {
				t.msStartUnix = int(time.Now().UnixMilli()) // время начала отсчета
				f.begin_flag = false
			}

			go func(f *flags_app) {

				for {
					time.Sleep(time.Millisecond * 50)
					if !f.paused {
						t.msTimer = int(time.Now().UnixMilli()) - t.msStartUnix // пройденное время
						t.ms = t.msTimer % 1000
						t.s = t.msTimer / 1000
						t.m = t.s / 60
						t.h = t.m / 60
						t.s = (t.msTimer - (t.h * 3600000) - (t.m * 60000) - t.ms) / 1000
						t.times = fmt.Sprintf("%02d:%02d:%02d:%03d", t.h, t.m, t.s, t.ms)
						w.input.Text = t.times
						w.input.Refresh() // Добавляем этот вызов для обновления виджета
					}

					if f.paused {
						break
					}

				}

			}(f)
		}

		updateButtonsState(f.start_pressed, f.paused)
	}

	// ПАУЗА
	pause := func() {
		f.start_pressed = false
		f.paused = true
		updateButtonsState(f.start_pressed, f.paused)
		//фиксируем время остановки в мс
		t.timePause = t.msTimer
	}

	// ОЧИСТКА ПОЛЯ
	reset := func() {
		w.input.Text = "00:00:00:000"
		w.input.Refresh()                           // Добавляем этот вызов для обновления виджета
		t.msStartUnix = int(time.Now().UnixMilli()) // это нужно если кнопка очистки нажата, но не нажата пауза. Тогда происходит очистка и продолжается счет
		f.begin_flag = true                         // на случай очищения на паузе (для запуска нового отсчета при нажатии старта)
		if f.paused {
			w.btn_clear.Disable()
		}

	}

	if s == "start" {
		return start
	} else if s == "pause" {
		return pause
	}

	// если вызывающий объект ничего не написал в параметре s то на выходе получит функцию стирания
	return reset

}
