package ui

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)

	app.GoPaintWindow.SetContent(swatchesContainer)
}
