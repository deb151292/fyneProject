package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyneCode.go/disp"
	"fyneCode.go/images"
	savedata "fyneCode.go/saveData"
)

type appInfo struct {
	name string
	icon fyne.Resource
	canv bool
	run  func(fyne.Window) fyne.CanvasObject
}

var apps = []appInfo{
	{"Save User", images.Saveicon, false, savedata.Insert},
	{"Display User", images.Dispicon, false, disp.User_Display},
}

func main() {

	a := app.New()
	a.SetIcon(images.ResourceIconPng)

	content := container.NewMax()
	w := a.NewWindow("Shop")

	appList := widget.NewList(
		func() int {
			return len(apps)
		},
		func() fyne.CanvasObject {
			icon := &canvas.Image{}
			label := widget.NewLabel("Text Editor")
			labelHeight := label.MinSize().Height
			icon.SetMinSize(fyne.NewSize(labelHeight, labelHeight))
			return container.NewBorder(nil, nil, icon, nil,
				label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			img := obj.(*fyne.Container).Objects[1].(*canvas.Image)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			img.Resource = apps[id].icon
			img.Refresh()
			text.SetText(apps[id].name)
		})
	appList.OnSelected = func(id widget.ListItemID) {
		content.Objects = []fyne.CanvasObject{apps[id].run(w)}

	}

	split := container.NewHSplit(appList, content)
	split.Offset = 0.2
	w.SetContent(split)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

//################################ FOR CREATING BYTE ARRAY THAT CAN BE USED IN USING ICONS IN FYNE ##################################################################
//for icon file

// func main() {
// 	iconFile, err := os.Open("C:/Users/deboj/OneDrive/Desktop/table/174314.png")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer iconFile.Close()

// 	// png.Decode() to convert image path to image
// 	img, err := png.Decode(iconFile)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Then convert the image to byte[] format using bytes.Buffer
// 	buf := new(bytes.Buffer)
// 	err1 := jpeg.Encode(buf, img, nil)
// 	_ = err1

// 	// save the converted byte[] to variable
// 	img_buf := buf.Bytes()

// 	s := make([]string, len(img_buf))
// 	for i := range img_buf {
// 		s[i] = strconv.Itoa(int(img_buf[i]))
// 	}
// 	log.Println(strings.Join(s, ","))
// }

//############################################################################################
