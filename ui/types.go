package ui

import (
	"go-paint/apptype"
	"go-paint/swatch"

	"fyne.io/fyne/v2"
)

type AppInit struct {
	GoPaintWindow fyne.Window
	State         *apptype.State
	Swatches      []*swatch.Swatch
}
