//go:build gui
// +build gui

package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("eDonish Pro - GUI")

	username := widget.NewEntry()
	username.SetPlaceHolder("Логин")

	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Пароль")

	loginBtn := widget.NewButton("Войти", func() {
		widget.NewLabel("Login clicked!")
	})

	content := container.NewVBox(
		widget.NewLabel("eDonish Pro"),
		username,
		password,
		loginBtn,
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}
