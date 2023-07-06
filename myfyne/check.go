package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type Check struct {
	widget.Check
}

func NewCheck(title string) *Check {
	b := &Check{}
	b.ExtendBaseWidget(b)
	b.Text = title
	return b
}
