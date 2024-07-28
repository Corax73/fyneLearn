package main

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type State struct {
	Val1, Val2        int
	Action            string
	IsAction, IsError bool
}

type Calc struct {
	State
	Input     *widget.Entry
	Display   *widget.Label
	CalcError string
}

func main() {
	myApp := app.New()
	c := Calc{Input: widget.NewEntry(), Display: widget.NewLabel(""), CalcError: "incorrect data, try again. please."}
	c.Input.SetPlaceHolder("Enter number")
	window := myApp.NewWindow("Calc")
	icon, err := fyne.LoadResourceFromPath("/Icon.png")
	if err == nil {
		window.SetIcon(icon)
	}
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter number")
	btnSum := widget.NewButton("+", func() {
		c.sumHandler(c.Input, c.Display)
	})
	btnSub := widget.NewButton("-", func() {
		c.subHandler(c.Input, c.Display)
	})

	btnEquals := widget.NewButton("=", func() {
		/**
		* @todo all action
		**/
		c.Val2, err = strconv.Atoi(c.Input.Text)
		var str string
		if err == nil {
			switch c.Action {
			case "+":
				res := c.Val1 + c.Val2
				str = strconv.Itoa(res)
			case "-":
				res := c.Val1 - c.Val2
				str = strconv.Itoa(res)
			default:
				str = ""
			}
		} else {
			str = c.CalcError
			c.IsError = true
			c.resetState()
		}
		c.resetState()
		c.Input.SetText(str)
	})

	btnClear := widget.NewButton("Clear", func() { c.clear() })

	btnExit := widget.NewButton("Exit", func() {
		myApp.Quit()
	})
	window.SetContent(
		container.NewGridWithColumns(
			2,
			c.Input,
			c.Display,
			container.NewGridWithColumns(4,
				c.addNumbBtn(1),
				c.addNumbBtn(2),
				c.addNumbBtn(3),
				c.addNumbBtn(4)),
			container.NewGridWithColumns(4,
				c.addNumbBtn(5),
				c.addNumbBtn(6),
				c.addNumbBtn(7),
				c.addNumbBtn(8)),
			container.NewGridWithColumns(4,
				c.addNumbBtn(9),
				c.addNumbBtn(0),
				btnSum,
				btnSub),
			container.NewGridWithColumns(3,
				btnEquals,
				btnClear,
				btnExit),
		),
	)
	window.Resize(fyne.NewSize(300, 200))
	window.ShowAndRun()
}

func (calc *Calc) subHandler(input *widget.Entry, display *widget.Label) {
	val, err := strconv.Atoi(input.Text)
	if err != nil {
		input.SetText(calc.CalcError)
		calc.IsError = true
		calc.resetState()
	} else {
		calc.IsAction = true
		calc.Action = "-"
		if calc.Val1 == 0 {
			calc.Val1 = val
			display.SetText("-")
		} else {
			calc.Val2 = val
			equal := calc.Val1 - calc.Val2
			res := strconv.Itoa(equal)
			input.SetText(res)
			display.SetText("")
			calc.Val1 = 0
			calc.Val2 = 0
		}
	}
}

func (calc *Calc) sumHandler(input *widget.Entry, display *widget.Label) {
	val, err := strconv.Atoi(input.Text)
	if err != nil {
		input.SetText(calc.CalcError)
		calc.IsError = true
		calc.resetState()
	} else {
		calc.IsAction = true
		calc.Action = "+"
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
		if val != "0" && !calc.IsAction && !calc.IsError {
			var strBuilder strings.Builder
			strBuilder.WriteString(val)
			strBuilder.WriteString(str)
			newVal = strBuilder.String()
			strBuilder.Reset()
		} else {
			calc.IsAction = false
			newVal = str
		}
		calc.IsError = false
		calc.Input.SetText(newVal)
	})
}

func (calc *Calc) resetState() {
	calc.Val1, calc.Val2 = 0, 0
	calc.IsAction = false
	calc.Action = ""
}

func (calc *Calc) clear() {
	calc.Input.SetText("0")
	calc.Display.SetText("")
	calc.resetState()
}
