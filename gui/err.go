package Gui

import (
	"../common"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"image/color"
	"os"
)

func ErrorAlert(msg string) {

	myApp := app.New()
	myWindow := myApp.NewWindow("linode-UI error")

	//img := canvas.NewImageFromResource(theme.CancelIcon())
	//img.FillMode = canvas.ImageFillOriginal
	//
	text := canvas.NewText("error:"+msg, color.White)
	content := fyne.NewContainerWithLayout(layout.NewCenterLayout(),
		text)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()

}

func InputToken(msg string) {

	myApp := app.New()
	myWindow := myApp.NewWindow("linode-UI error")

	//img := canvas.NewImageFromResource(theme.CancelIcon())
	//img.FillMode = canvas.ImageFillOriginal
	//

	f := widget.NewForm()
	f.Append("err", widget.NewLabel(msg))
	_password := widget.NewPasswordEntry()
	f.Append("REINPUT token?", _password)
	f.Append("action", widget.NewButton("SAVE&QUIT linode-UI", func() {
		if len(_password.Text) == 64 {
			Common.Config.Set("token", _password.Text)
			Common.Config.Set("timeout", Common.Config.Get("timeout"))
			Common.Config.Set("retryTime", Common.Config.Get("retryTime"))

			err = Common.Config.WriteConfig()

		}
		os.Exit(2)
	}))

	content := fyne.NewContainerWithLayout(layout.NewFormLayout(),
		f)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 200))
	myWindow.ShowAndRun()

}
