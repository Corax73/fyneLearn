package customTheme

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type CustomTheme struct {
	fyne.Theme
}

func NewCustomTheme() fyne.Theme {
	return &CustomTheme{Theme: theme.DefaultTheme()}
}

func (cTheme *CustomTheme) Color(name fyne.ThemeColorName, _ fyne.ThemeVariant) color.Color {
	return cTheme.Theme.Color(name, theme.VariantDark)
}
