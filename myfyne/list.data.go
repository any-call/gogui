package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	ListData struct {
		widget.List
	}

	ListLength     func() int
	ListCreateCell func() fyne.CanvasObject
	ListUpdateCell func(id widget.ListItemID, c fyne.CanvasObject)
)

func NewListData(length ListLength, createItem ListCreateCell, updateItem ListUpdateCell) *ListData {
	w := &ListData{}
	w.ExtendBaseWidget(w)
	w.Length = length
	w.CreateItem = createItem
	w.UpdateItem = updateItem
	return w
}
