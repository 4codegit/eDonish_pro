package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("eDonish Pro")

	content := container.NewVBox(
		widget.NewLabel("eDonish Pro"),
		widget.NewButton("Click", func() {}),
	)

	w.SetContent(content)
	w.Resize(fyne.NewSize(400, 300))
	w.ShowAndRun()
}

