package home

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyneCode.go/images"
)

func Home(win fyne.Window) fyne.CanvasObject {

	return container.NewBorder(nil, nil, nil, nil, canvas.NewImageFromResource(images.Home))
}
