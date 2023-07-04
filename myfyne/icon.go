package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/gobase/frame/mysignal"
)

type Icon struct {
	widget.Icon
	OnTapped       mysignal.Signal[*fyne.PointEvent]
	OnDoubleTapped mysignal.Signal[*fyne.PointEvent]
}

func NewIcon(res fyne.Resource) *Icon {
	if res == nil {
		return nil
	}

	icon := &Icon{}
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)
	return icon
}

func (self *Icon) Tapped(event *fyne.PointEvent) {
	self.OnTapped.AsyncEmit(event)
}

func (self *Icon) DoubleTapped(event *fyne.PointEvent) {
	self.OnDoubleTapped.AsyncEmit(event)
}
