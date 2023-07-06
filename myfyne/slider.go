package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type Slider struct {
	widget.Slider
}

func NewSlider(min, max float64) *Slider {
	b := &Slider{}
	b.ExtendBaseWidget(b)
	b.Min = min
	b.Max = max
	return b
}
