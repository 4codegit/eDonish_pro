package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("eDonish Pro")
	
	username := widget.NewEntry()
	username.SetPlaceHolder("Логин")
	
	password := widget.NewPasswordEntry()
	password.SetPlaceHolder("Пароль")
	
	loginBtn := widget.NewButton("Войти", func() {})
	
	content := fyne.NewContainerWithLayout(fyne.NewVBox(
		widget.NewLabel("eDonish Pro"),
		username,
		password,
		loginBtn,
	))
	
	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

