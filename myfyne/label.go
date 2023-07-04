package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Label struct {
	widget.Label
}

func NewLabel(txt string) *Label {
	label := &Label{}
	label.ExtendBaseWidget(label)
	label.SetText(txt)

	return label
}

func (self *Label) Color(c fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(c, v)
}

func (self *Label) TextColor() color.Color {
	return color.Black
}
