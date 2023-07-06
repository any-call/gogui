package myfyne

import (
	"fyne.io/fyne/v2/widget"
)

type TextView struct {
	widget.RichText
}

func NewTextView(text string) *TextView {
	b := &TextView{}
	b.ExtendBaseWidget(b)
	b.Segments = []widget.RichTextSegment{
		&widget.TextSegment{Text: text, Style: widget.RichTextStyleInline},
	}

	return b
}
