package myfyne

import (
	"fyne.io/fyne/v2/widget"
	"net/url"
)

type HyperLink struct {
	widget.Hyperlink
}

func NewHyperLink(text string, url *url.URL) *HyperLink {
	b := &HyperLink{}
	b.ExtendBaseWidget(b)
	b.SetText(text)
	if url != nil {
		b.SetURL(url)
	}
	return b
}
