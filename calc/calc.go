package calc

import (
	"strconv"
	"strings"

	"fyne.io/fyne/v2/widget"
)

type State struct {
	Val1, Val2                  float64
	Action                      string
	IsAction, IsError, IsResult bool
}

type Calc struct {
	State
	Input     *widget.Entry
	Display   *widget.Label
	CalcError string
}

func (calc *Calc) SubHandler(input *widget.Entry, display *widget.Label) {
	val, err := strconv.ParseFloat(input.Text, 64)
	if err != nil {
		input.SetText(calc.CalcError)
		calc.IsError = true
		calc.ResetState()
	} else {
		calc.IsAction = true
		calc.Action = "-"
		if calc.Val1 == 0 {
			calc.Val1 = val
			display.SetText("-")
		} else {
			calc.Val2 = val
			equal := calc.Val1 - calc.Val2
			res := strconv.FormatFloat(equal, 'f', 2, 64)
			input.SetText(res)
			display.SetText("")
			calc.Val1 = 0
			calc.Val2 = 0
		}
	}
}

func (calc *Calc) SumHandler(input *widget.Entry, display *widget.Label) {
	val, err := strconv.ParseFloat(input.Text, 64)
	if err != nil {
		input.SetText(calc.CalcError)
		calc.IsError = true
		calc.ResetState()
	} else {
		calc.IsAction = true
		calc.Action = "+"
		if calc.Val1 == 0 {
			calc.Val1 = val
			display.SetText("+")
		} else {
			calc.Val2 = val
			equal := calc.Val1 + calc.Val2
			res := strconv.FormatFloat(equal, 'f', 2, 64)
			input.SetText(res)
			display.SetText("")
			calc.Val1 = 0
			calc.Val2 = 0
		}
	}
}

func (calc *Calc) DivHandler(input *widget.Entry, display *widget.Label) {
	val, err := strconv.ParseFloat(input.Text, 64)
	if err != nil {
		input.SetText(calc.CalcError)
		calc.IsError = true
		calc.ResetState()
	} else {
		calc.IsAction = true
		calc.Action = "/"
		if calc.Val1 == 0 {
			calc.Val1 = val
			display.SetText("/")
		} else {
			calc.Val2 = val
			equal := calc.Val1 / calc.Val2
			res := strconv.FormatFloat(equal, 'f', 2, 64)
			input.SetText(res)
			display.SetText("")
			calc.Val1 = 0
			calc.Val2 = 0
		}
	}
}

func (calc *Calc) MultHandler(input *widget.Entry, display *widget.Label) {
	val, err := strconv.ParseFloat(input.Text, 64)
	if err != nil {
		input.SetText(calc.CalcError)
		calc.IsError = true
		calc.ResetState()
	} else {
		calc.IsAction = true
		calc.Action = "*"
		if calc.Val1 == 0 {
			calc.Val1 = val
			display.SetText("*")
		} else {
			if val != 0 {
				calc.Val2 = val
				equal := calc.Val1 * calc.Val2
				res := strconv.FormatFloat(equal, 'f', 2, 64)
				input.SetText(res)
				display.SetText("")
				calc.Val1 = 0
				calc.Val2 = 0
			} else {
				calc.IsError = true
				display.SetText("Division by zero")
			}
		}
	}
}

func (calc *Calc) AddNumbBtn(number int) *widget.Button {
	str := strconv.Itoa(number)
	return widget.NewButton(str, func() {
		val := calc.Input.Text
		var newVal string
		if val != "0" && !calc.IsAction && !calc.IsError && !calc.IsResult {
			var strBuilder strings.Builder
			strBuilder.WriteString(val)
			strBuilder.WriteString(str)
			newVal = strBuilder.String()
			strBuilder.Reset()
		} else {
			calc.IsAction = false
			calc.IsResult = false
			newVal = str
		}
		calc.IsError = false
		calc.Input.SetText(newVal)
	})
}

func (calc *Calc) ResetState() {
	calc.Val1, calc.Val2 = 0, 0
	calc.IsAction = false
	calc.Action = ""
}

func (calc *Calc) Clear() {
	calc.Input.SetText("0")
	calc.Display.SetText("")
	calc.ResetState()
}
