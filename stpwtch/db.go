package stopwatch

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func begin_work_db() {
	// инициализация видов активностей
	var names_activites []string

	result := inf_app.DB.Model(&Activity{}).Pluck("Name_activity", &names_activites)
	if result.Error != nil {
		log.Fatal("Ошибка при получении имен:", result.Error)
	}

	wdgts.ent_activity = widget.NewSelect(
		names_activites,

		func(s string) {
			fmt.Printf("Selected is %s", s)
		},
	)
	wdgts.ent_activity.PlaceHolder = "Выберите активность"
	wdgts.ent_activity.Hide()

	//срез видов активностей
	var activites []Activity

	//БЛОК ИНИЦИАЛИЗАЦИИ БАЗЫ ДАННЫХ
	var err error
	inf_app.DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}
	inf_app.DB.AutoMigrate(&Activity{})
	inf_app.DB.Find(&activites)

}

func adding_an_activity() {
	added_act := Activity{
		Name_activity: wdgts.input_act.Text,
		Description:   wdgts.input_act_descr.Text,
	}

	inf_app.DB.Create(&added_act)
	inf_app.DB.Find(&added_act)

	// после добавления активности - ворачиваем контент секундомера
	wdgts.w.SetContent(wdgts.cont_begin)                   // восстановление первоначального вида секундомера
	wdgts.w.Resize(fyne.NewSize(inf_app.w_a, inf_app.h_a)) // восстановление размера окна
}
