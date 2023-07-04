package myfyne

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func NewApp() fyne.App {
	return app.New()
}

func NewAppWithID(id string) fyne.App {
	return app.NewWithID(id)
}

func test() {
	//NewApp().NewWindow().setsi
	//NewApp().NewWindow("wii").
}
