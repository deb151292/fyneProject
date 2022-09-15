package savedata

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func Insert(win fyne.Window) fyne.CanvasObject {

	var form *widget.Form

	UserName := widget.NewEntry()

	// textArea := widget.NewMultiLineEntry()

	Email := widget.NewEntry()

	// textArea1 := widget.NewMultiLineEntry()
	Contact := widget.NewEntry()

	form = &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "User", Widget: UserName}, {Text: "Email", Widget: Email}, {Text: "Contact Number", Widget: Contact}},
		OnSubmit: func() { // optional, handle form submission

			massage := User_insert_json(UserName.Text, Email.Text, Contact.Text)

			dialog.ShowInformation("Massage", massage, fyne.CurrentApp().Driver().AllWindows()[0])

			UserName.SetText("")
			Email.SetText("")
			Contact.SetText("")

		}, OnCancel: func() {
			UserName.SetText("")
			Email.SetText("")
			Contact.SetText("")
		},
	}

	// we can also append items
	// form.Append(textArea)
	// form.Append("Text1", textArea1)

	return container.NewBorder(nil, nil, nil, nil, form)

}
