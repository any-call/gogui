package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	ListTree struct {
		widget.Tree
	}

	TreeChilds     func(uid string) []string
	TreeIsBranch   func(uid string) bool
	TreeCreateCell func(branch bool) fyne.CanvasObject
	TreeUpdateCell func(uid string, branch bool, node fyne.CanvasObject)
)

func NewListTree(childs TreeChilds, isBranch TreeIsBranch,
	createCell TreeCreateCell,
	updateCell TreeUpdateCell) *ListTree {
	w := &ListTree{}
	w.ExtendBaseWidget(w)
	w.ChildUIDs = childs
	w.IsBranch = isBranch
	w.CreateNode = createCell
	w.UpdateNode = updateCell
	return w
}
