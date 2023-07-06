package myfyne

import "fyne.io/fyne/v2/widget"

type Toolbar struct {
	widget.Toolbar
}

func NewToolbar() *Toolbar {
	w := &Toolbar{}
	w.ExtendBaseWidget(w)
	return w
}
