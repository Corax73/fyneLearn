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
	Input   *widget.Entry
	Display *widget.Label
}

func main() {
	myApp := app.New()
	c := Calc{Input: widget.NewEntry(), Display: widget.NewLabel("")}
	c.Input.SetPlaceHolder("Enter number")
	window := myApp.NewWindow("Calc")
	icon, err := fyne.LoadResourceFromPath("/Icon.png")
	if err == nil {
		window.SetIcon(icon)
	}
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter number")
	btnSum := widget.NewButton("+", func() {
		sumHandler(&c, c.Input, c.Display)
	})
	btnSub := widget.NewButton("-", func() {
		val1, err := strconv.Atoi(c.Input.Text)
		c.Input.SetText("0")
		if err != nil {
			fmt.Println(err)
		}
		val2, err := strconv.Atoi(c.Display.Text)
		if err != nil {
			fmt.Println(err)
		}
		if val2 == 0 {
			val2 = val1 * 2
		}
		res := strconv.Itoa(sub(val2, val1))
		c.Display.SetText(res)
	})

	btnEquals := widget.NewButton("=", func() {
		/**
		* @todo all action
		**/
		c.Val2, err = strconv.Atoi(c.Input.Text)
		var str string
		if err == nil {
			res := c.Val1 + c.Val2
			str = strconv.Itoa(res)
		} else {
			str = err.Error()
		}
		c.Input.SetText(str)
	})

	btnClear := widget.NewButton("Clear", func() {
		c.Input.SetText("0")
	})

	btnExit := widget.NewButton("Exit", func() {
		myApp.Quit()
	})
	window.SetContent(
		container.NewVBox(
			c.Input,
			c.addNumbBtn(1),
			c.addNumbBtn(2),
			c.addNumbBtn(3),
			c.addNumbBtn(4),
			c.addNumbBtn(5),
			c.addNumbBtn(6),
			c.addNumbBtn(7),
			c.addNumbBtn(8),
			c.addNumbBtn(9),
			c.addNumbBtn(0),
			btnSum,
			btnSub,
			btnEquals,
			c.Display,
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

func sumHandler(calc *Calc, input *widget.Entry, display *widget.Label) {
	val, err := strconv.Atoi(input.Text)
	if err != nil {
		input.SetText(err.Error())
	} else {
		if calc.Val1 == 0 {
			calc.Val1 = val
			display.SetText("+")
		} else {
			calc.Val2 = val
			equal := calc.Val1 + calc.Val2
			res := strconv.Itoa(equal)
			input.SetText(res)
			display.SetText("")
			calc.Val1 = 0
			calc.Val2 = 0
		}
	}
}

func (calc *Calc) addNumbBtn(number int) *widget.Button {
	str := strconv.Itoa(number)
	return widget.NewButton(str, func() {
		val := calc.Input.Text
		var newVal string
		if val != "0" {
			var strBuilder strings.Builder
			strBuilder.WriteString(val)
			strBuilder.WriteString(str)
			newVal = strBuilder.String()
			strBuilder.Reset()
		} else {
			newVal = str
		}
		calc.Input.SetText(newVal)
	})
}
