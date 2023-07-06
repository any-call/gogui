package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type Button struct {
	widget.Button
}

func NewButton(title string) *Button {
	b := &Button{}
	b.ExtendBaseWidget(b)
	b.SetText(title)
	return b
}
