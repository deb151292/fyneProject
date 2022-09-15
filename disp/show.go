package disp

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"fyneCode.go/entity"
)

func User_Display(win fyne.Window) fyne.CanvasObject {
	OldUser := []entity.User{}
	input, err := ioutil.ReadFile("userData.json")
	json.Unmarshal(input, &OldUser)
	if err != nil {
		log.Println("error")
	}

	data := make([][]string, len(OldUser)+1)

	n := len(OldUser) + 1
	rows := make([]string, n*3)
	for i := 0; i < n; i++ {
		data[i] = rows[i*3 : (i+1)*3]
	}

	data[0][0] = "User Name"
	data[0][1] = "Email"
	data[0][2] = "Phone Number"

	for i, v := range OldUser {
		data[i+1][0] = v.UserName
		data[i+1][1] = v.Email
		data[i+1][2] = v.PhoneNumber
	}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("User")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})
	list.SetColumnWidth(0, 300)
	list.SetColumnWidth(1, 300)
	list.SetColumnWidth(2, 300)

	return container.NewBorder(nil, nil, nil, nil, list)

}
