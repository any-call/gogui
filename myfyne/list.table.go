package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type (
	ListTable struct {
		widget.Table
	}

	TableLength     func() (int, int)
	TableCreateCell func() fyne.CanvasObject
	TableUpdateCell func(id widget.TableCellID, c fyne.CanvasObject)
)

func NewListTable(len TableLength, createItem TableCreateCell, updateItem TableUpdateCell) *ListTable {
	w := &ListTable{}
	w.ExtendBaseWidget(w)
	w.Length = len
	w.CreateCell = createItem
	w.UpdateCell = updateItem
	return w
}
