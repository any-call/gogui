package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type (
	TextGrid struct {
		widget.TextGrid
	}
)

func NewTextGrid() *TextGrid {
	w := &TextGrid{}
	w.ExtendBaseWidget(w)
	return w
}
