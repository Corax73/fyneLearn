package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Learn")
	label := widget.NewLabel("Hello")
	window.SetContent(
		container.NewVBox(
			label,
			widget.NewButton("Exit", func() {
				myApp.Quit()
			})),
	)
	window.Resize(fyne.NewSize(540, 480))
	window.ShowAndRun()
}
