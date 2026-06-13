package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("eDonish Pro")
	w.SetContent(widget.NewLabel("Hello eDonish Pro!"))
	w.ShowAndRun()
}

