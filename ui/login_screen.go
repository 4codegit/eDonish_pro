package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/container"
)

func NewLoginScreen() fyne.CanvasObject {
	username := widget.NewEntry()
	username.SetPlaceHolder("Логин")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Пароль")

	loginBtn := widget.NewButton("Войти", func() {
		// TODO: Реализовать авторизацию
	})

	return container.NewVBox(
		widget.NewLabel("eDonish Pro - Авторизация"),
		widget.NewForm(
			widget.NewFormItem("Логин", username),
			widget.NewFormItem("Пароль", password),
		),
		loginBtn,
	)
}

