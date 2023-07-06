package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TextField struct {
	widget.Entry
}

func NewTextField() *TextField {
	b := &TextField{}
	b.ExtendBaseWidget(b)
	b.Wrapping = fyne.TextTruncate
	return b
}

func (tf *TextField) SetPassword(b bool) *TextField {
	tf.Password = b
	return tf
}

func (tf *TextField) SetValidor(cb fyne.StringValidator) *TextField {
	tf.Validator = cb
	return tf
}
