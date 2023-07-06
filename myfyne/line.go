package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

type Line struct {
	widget.BaseWidget
	col color.Color
}

func NewLine(c color.Color) *Line {
	s := &Line{}
	s.ExtendBaseWidget(s)
	if c == nil {
		s.col = theme.SeparatorColor()
	} else {
		s.col = c
	}

	s.col = c
	return s
}

func (s *Line) MinSize() fyne.Size {
	s.ExtendBaseWidget(s)
	t := theme.SeparatorThicknessSize()
	return fyne.NewSize(t, t)
}

// Implements: fyne.Widget
func (s *Line) CreateRenderer() fyne.WidgetRenderer {
	s.ExtendBaseWidget(s)
	bar := canvas.NewRectangle(theme.SeparatorColor())
	return &lineRenderer{
		WidgetRenderer: widget.NewSimpleRenderer(bar),
		bar:            bar,
		d:              s,
	}
}

type lineRenderer struct {
	fyne.WidgetRenderer
	bar *canvas.Rectangle
	d   *Line
}

func (r *lineRenderer) MinSize() fyne.Size {
	t := theme.SeparatorThicknessSize()
	return fyne.NewSize(t, t)
}

func (r *lineRenderer) Refresh() {
	r.bar.FillColor = r.d.col
	canvas.Refresh(r.d)
}
