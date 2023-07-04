package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

type Button struct {
	BaseWidget
	Text string

	background      color.Color
	hoverBackground color.Color
	textColor       color.Color
	highColor       color.Color
	//event
	OnTapped func() `json:"-"`
}

func NewButton(title string, tapped func()) *Button {
	button := &Button{
		Text:            title,
		OnTapped:        tapped,
		background:      theme.BackgroundColor(),
		hoverBackground: color.Gray{Y: 100},
		textColor:       theme.TextColor(),
		highColor:       color.Transparent,
	}

	button.ExtendBaseWidget(button)
	return button
}

func (b *Button) CreateRenderer() fyne.WidgetRenderer {
	b.ExtendBaseWidget(b)
	background := canvas.NewRectangle(b.hoverBackground)

	label := canvas.NewText(b.Text, b.textColor)
	r := &buttonReader{
		background: background,
		label:      label,
		button:     b,
	}

	return r
}

type buttonReader struct {
	button     *Button
	background *canvas.Rectangle
	label      *canvas.Text
}

// 该接口，用于布局 componse
func (r *buttonReader) Layout(Size fyne.Size) {
	if r.button.hovered {
		r.background = canvas.NewRectangle(r.button.hoverBackground)
	} else {
		r.background = canvas.NewRectangle(r.button.background)
	}

	r.background.Move(fyne.NewPos(0, 0))
	r.background.Resize(Size)
	if r.label.Text != "" {
		r.label.Move(fyne.NewPos(0, 0))
		r.label.Resize(r.label.MinSize())
	}
}

func (r *buttonReader) MinSize() fyne.Size {
	return fyne.NewSize(100, 100)
}

func (r *buttonReader) Refresh() {
	canvas.Refresh(r.button)
}

func (r *buttonReader) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.background, r.label}
}

func (r *buttonReader) Destroy() {
	r.background = nil
	r.label = nil
}
