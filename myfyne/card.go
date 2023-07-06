package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type Card struct {
	widget.Card
}

func NewCard(title, subTitle string, co fyne.CanvasObject) *Card {
	w := &Card{}
	w.ExtendBaseWidget(w)

	w.SetTitle(title)
	w.SetSubTitle(subTitle)
	if co != nil {
		w.SetContent(co)
	}

	return w
}
