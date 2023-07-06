package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type ComboField struct {
	widget.SelectEntry
}

func NewComboField(option []string) *ComboField {
	b := &ComboField{}
	b.ExtendBaseWidget(b)
	b.SetOptions(option)
	return b
}
