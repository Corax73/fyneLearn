package main

import (
	"learnFyne/calc"
	"learnFyne/customTheme"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	c := calc.Calc{Input: widget.NewEntry(), Display: widget.NewLabel(""), CalcError: "incorrect data, try again. please."}
	myApp.Settings().SetTheme(customTheme.NewCustomTheme())
	c.Input.SetPlaceHolder("Enter number")
	window := myApp.NewWindow("Calc")
	icon, err := fyne.LoadResourceFromPath("/Icon.png")
	if err == nil {
		window.SetIcon(icon)
	}
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter number")

	btnSum := widget.NewButton("+", func() {
		c.SumHandler(c.Input, c.Display)
	})

	btnSub := widget.NewButton("-", func() {
		c.SubHandler(c.Input, c.Display)
	})

	btnDiv := widget.NewButton("/", func() {
		c.DivHandler(c.Input, c.Display)
	})

	btnMult := widget.NewButton("*", func() {
		c.MultHandler(c.Input, c.Display)
	})

	btnEquals := widget.NewButton("=", func() {
		c.Val2, err = strconv.ParseFloat(c.Input.Text, 64)
		var str string
		if err == nil {
			switch c.Action {
			case "+":
				res := c.Val1 + c.Val2
				c.IsAction = true
				str = strconv.FormatFloat(res, 'f', 2, 64)
			case "-":
				res := c.Val1 - c.Val2
				c.IsAction = true
				str = strconv.FormatFloat(res, 'f', 2, 64)
			case "/":
				if c.Val2 != 0 {
					res := c.Val1 / c.Val2
					str = strconv.FormatFloat(res, 'f', 2, 64)
				} else {
					c.IsError = true
					c.Display.SetText("Division by zero")
					c.ResetState()
				}
				c.IsAction = true
			case "*":
				res := c.Val1 * c.Val2
				c.IsAction = true
				str = strconv.FormatFloat(res, 'f', 2, 64)
			default:
				str = ""
			}
		} else {
			str = c.CalcError
			c.IsError = true
			c.ResetState()
		}
		c.ResetState()
		c.IsResult = true
		c.Display.SetText("")
		c.Input.SetText(str)
	})

	btnClear := widget.NewButton("Clear", func() { c.Clear() })

	btnExit := widget.NewButton("Exit", func() {
		myApp.Quit()
	})

	window.SetContent(
		container.NewGridWithColumns(
			1,
			c.Input,
			c.Display,
			container.NewGridWithColumns(3,
				c.AddNumbBtn(1),
				c.AddNumbBtn(2),
				c.AddNumbBtn(3)),
			container.NewGridWithColumns(3,
				c.AddNumbBtn(4),
				c.AddNumbBtn(5),
				c.AddNumbBtn(6)),
			container.NewGridWithColumns(3,
				c.AddNumbBtn(7),
				c.AddNumbBtn(8),
				c.AddNumbBtn(9)),
			container.NewGridWithColumns(1,
				c.AddNumbBtn(0)),
			container.NewGridWithColumns(3,
				btnSum,
				btnSub,
				btnDiv),
			container.NewGridWithColumns(3,
				btnMult,
				btnEquals,
				btnClear),
			container.NewGridWithColumns(1,
				btnExit),
		),
	)
	window.Resize(fyne.NewSize(300, 200))
	window.ShowAndRun()
}
