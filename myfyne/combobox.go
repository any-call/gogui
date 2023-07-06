package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type ComboBox struct {
	widget.Select
}

func NewComboBox(option []string) *ComboBox {
	b := &ComboBox{}
	b.ExtendBaseWidget(b)
	b.Options = option

	return b
}
