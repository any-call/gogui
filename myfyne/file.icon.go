package myfyne

import (
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/any-call/gobase/util/myos"
)

type FileIcon struct {
	widget.FileIcon
}

func NewFileIcon(path string) *FileIcon {
	b := &FileIcon{}
	b.ExtendBaseWidget(b)
	if myos.IsExistFile(path) {
		b.SetURI(storage.NewFileURI(path))
	}

	return b
}
