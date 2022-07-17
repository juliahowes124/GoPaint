package main

import (
	"go-paint/apptype"
	"go-paint/swatch"
	"go-paint/ui"
	"image/color"

	"fyne.io/fyne/v2/app"
)

func main() {
	goPaintApp := app.New()
	goPaintWindow := goPaintApp.NewWindow("go-paint")
	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}
	appInit := ui.AppInit{
		GoPaintWindow: goPaintWindow,
		State:         &state,
		Swatches:      make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)
	appInit.GoPaintWindow.ShowAndRun()
}
