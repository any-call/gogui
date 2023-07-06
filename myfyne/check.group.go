package myfyne

import "fyne.io/fyne/v2/widget"

type CheckGroup struct {
	widget.CheckGroup
}

func NewCheckGroup(option []string) *CheckGroup {
	b := &CheckGroup{}
	b.ExtendBaseWidget(b)
	b.Options = option
	return b
}
