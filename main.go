package main

import (
	"fmt"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type State struct {
	Val1, Val2 int
}

type Calc struct {
	State
	Display *widget.Label
}

func main() {
	myApp := app.New()
	var data State
	window := myApp.NewWindow("Calc")
	icon, err := fyne.LoadResourceFromPath("/Icon.png")
	if err == nil {
		window.SetIcon(icon)
	}
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter number")
	display := widget.NewLabel("0")
	btnSum := widget.NewButton("+", func() {
		sumHandler(&data, input, display)
	})
	btnSub := widget.NewButton("-", func() {
		val1, err := strconv.Atoi(input.Text)
		input.SetText("0")
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
			input,
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

func sub(val2, val1 int) int {
	return val2 - val1
}

func sumHandler(data *State, input *widget.Entry, display *widget.Label) {
	val, err := strconv.Atoi(input.Text)
	if err != nil {
		input.SetText(err.Error())
	} else {
		if data.Val1 == 0 {
			data.Val1 = val
			display.SetText("+")
		} else {
			data.Val2 = val
			equal := data.Val1 + data.Val2
			res := strconv.Itoa(equal)
			input.SetText(res)
			display.SetText("")
			data.Val1 = 0
			data.Val2 = 0
		}
	}
}

func (calc *Calc) addNumbBtn(number int) *widget.Button {
	str := strconv.Itoa(number)
	return widget.NewButton(str, func() {
		val := calc.Display.Text
		var strBuilder strings.Builder
		strBuilder.WriteString(val)
		strBuilder.WriteString(str)
		calc.Display.SetText(strBuilder.String())
		strBuilder.Reset()
	})
}
