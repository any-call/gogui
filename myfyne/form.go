package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type Form struct {
	widget.Form
}

func NewForm() *Form {
	b := &Form{}
	b.ExtendBaseWidget(b)
	return b
}
