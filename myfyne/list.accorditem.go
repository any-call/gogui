package myfyne

import "fyne.io/fyne/v2/widget"

type ListAccItem struct {
	widget.Accordion
}

func NewListAccItem() *ListAccItem {
	w := &ListAccItem{}
	w.ExtendBaseWidget(w)

	return w
}
