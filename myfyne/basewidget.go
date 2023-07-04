package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/gobase/frame/mysignal"
)

type BaseWidget struct {
	widget.BaseWidget
	cursor  desktop.Cursor
	focused bool
	hovered bool

	OnMouseIn    mysignal.Signal[*desktop.MouseEvent]
	OnMouseMoved mysignal.Signal[*desktop.MouseEvent]
	OnMouseOut   mysignal.Signal[any]
	OnTapped     mysignal.Signal[*fyne.PointEvent]
	OnTypedRune  mysignal.Signal[rune]
	OnTypedKey   mysignal.Signal[*fyne.KeyEvent]
}

func (b *BaseWidget) Cursor() desktop.Cursor {
	if b.cursor == nil {
		b.cursor = desktop.DefaultCursor
	}

	return b.cursor
}

func (b *BaseWidget) FocusGained() {
	b.focused = true
	b.Refresh()
}

func (b *BaseWidget) FocusLost() {
	b.focused = false
	b.Refresh()
}

func (b *BaseWidget) MouseIn(event *desktop.MouseEvent) {
	b.OnMouseIn.AsyncEmit(event)
	b.hovered = true
	b.Refresh()
}

func (b *BaseWidget) MouseMoved(event *desktop.MouseEvent) {
	b.OnMouseMoved.AsyncEmit(event)
}

func (b *BaseWidget) MouseOut() {
	b.OnMouseOut.AsyncEmit(nil)
	b.hovered = false
	b.Refresh()
}

func (b *BaseWidget) Tapped(event *fyne.PointEvent) {
	b.OnTapped.AsyncEmit(event)
}

func (b *BaseWidget) TypedRune(r rune) {
	b.OnTypedRune.AsyncEmit(r)
}

func (b *BaseWidget) TypedKey(event *fyne.KeyEvent) {
	b.OnTypedKey.AsyncEmit(event)
}
