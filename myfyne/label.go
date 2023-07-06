package myfyne

import (
	"fyne.io/fyne/v2/widget"
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
