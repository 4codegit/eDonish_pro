package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2"
)

func main() {
	a := app.New()
	w := a.NewWindow("eDonish Pro")
	w.SetContent(NewLoginScreen())
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

