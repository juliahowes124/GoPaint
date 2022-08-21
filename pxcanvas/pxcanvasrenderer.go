package pxcanvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type PxCanvasRenderer struct {
	pxCanvas     *PxCanvas     //parent
	canvasImage  *canvas.Image //part of fyne, display image defined in PxCanvas structure
	canvasBorder []canvas.Line
}

// Implementing WidgetRenderer

func (renderer *PxCanvasRenderer) MinSize() fyne.Size {
	return renderer.pxCanvas.DrawingArea
}

func (renderer *PxCanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5) // 4 are border lines, 1 is the image itself
	for i := 0; i < len(renderer.canvasBorder); i++ {
		objects = append(objects, &renderer.canvasBorder[i])
	}
	objects = append(objects, renderer.canvasImage)
	return objects
}

func (renderer *PxCanvasRenderer) Destroy() {}

func (renderer *PxCanvasRenderer) Layout(size fyne.Size) {}
