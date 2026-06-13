package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("eDonish Pro")

	username := widget.NewEntry()
	username.SetPlaceHolder("Логин")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Пароль")

	loginBtn := widget.NewButton("Войти", func() {
		if username.Text == "" || password.Text == "" {
			dialog.ShowError("Ошибка", "Заполните все поля", w)
			return
		}
		dialog.ShowInformation("Успех", "Вход выполнен!", w)
	})

	form := container.NewVBox(
		widget.NewLabel("eDonish Pro - Авторизация"),
		widget.NewSeparator(),
		username,
		password,
		loginBtn,
	)

	w.SetContent(form)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

