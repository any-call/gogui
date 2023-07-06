package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type Icon struct {
	widget.Icon
}

func NewIcon() *Icon {
	icon := &Icon{}
	icon.ExtendBaseWidget(icon)
	return icon
}
