package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.NewWithID("de.nroo.test")
	a.Settings().SetTheme(theme.LightTheme())
	w := a.NewWindow("Test")
	w.Resize(fyne.Size{Width: 640, Height: 380})

	about := fyne.NewContainerWithLayout(
		layout.NewCenterLayout(),
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			widget.NewLabelWithStyle("Testin asdf", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		),
	)

	content := widget.NewTabContainer(
		widget.NewTabItemWithIcon("About", theme.InfoIcon(), about),
	)

	content.SetTabLocation(widget.TabLocationLeading)
	w.SetContent(content)
	w.SetMaster()
	w.ShowAndRun()
}
