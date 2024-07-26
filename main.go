package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Calc")
	icon, err := fyne.LoadResourceFromPath("/Icon.png")
	if err == nil {
		window.SetIcon(icon)
	}
	input1 := widget.NewEntry()
	input1.SetPlaceHolder("Enter number")
	display := widget.NewLabel("0")
	btnSum := widget.NewButton("+", func() {
		val1, err := strconv.Atoi(input1.Text)
		input1.SetText("0")
		if err != nil {
			fmt.Println(err)
		}
		val2, err := strconv.Atoi(display.Text)
		if err != nil {
			fmt.Println(err)
		}
		res := strconv.Itoa(sum(val1, val2))
		display.SetText(res)
	})

	btnSub := widget.NewButton("-", func() {
		val1, err := strconv.Atoi(input1.Text)
		input1.SetText("0")
		if err != nil {
			fmt.Println(err)
		}
		val2, err := strconv.Atoi(display.Text)
		if err != nil {
			fmt.Println(err)
		}
		if val2 == 0 {
			val2 = val1 * 2
		}
		res := strconv.Itoa(sub(val2, val1))
		display.SetText(res)
	})

	btnClear := widget.NewButton("Clear", func() {
		display.SetText("0")
	})

	btnExit := widget.NewButton("Exit", func() {
		myApp.Quit()
	})
	window.SetContent(
		container.NewVBox(
			input1,
			btnSum,
			btnSub,
			display,
			btnClear,
			btnExit,
		),
	)
	window.Resize(fyne.NewSize(640, 180))
	window.ShowAndRun()
}

func sum(numbers ...int) int {
	var resp int
	for _, number := range numbers {
		resp += number
	}
	return resp
}

func sub(val2, val1 int) int {
	return val2 - val1
}
