package Gui

import (
	"../common"
	"../linodeApi"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"log"
	"net/url"
	"os"
	"runtime"
	"strconv"
)

var err error
var textLog = "loading....."

func Run() {

	if "darwin" == runtime.GOOS {
		os.Setenv("FYNE_FONT", "/System/Library/Fonts/STHeiti Medium.ttc")
	} else {
		os.Setenv("FYNE_FONT", "C:\\windows\\Fonts\\simhei.ttf")
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("Linode-UI")

	setTabContent(0, myWindow)

	myWindow.ShowAndRun()
}

func setTabContent(id int, window fyne.Window) {
	textLog = "loading....."
	displayFullScreenLog(displayLogNew("loading"), window)
	//label_1 := widget.NewLabel("Linode VPS Controller\n https://https://github.com/terry2010/linode-ui")
	u, _ := url.Parse("https://https://github.com/terry2010/linode-ui")

	cForm := widget.NewForm(widget.NewFormItem("Create Demo Linode", widget.NewLabel("Create JP/CentOS8/512M/25G Linode Now")))
	_label := widget.NewEntry()
	cForm.Append("label", _label)
	_password := widget.NewPasswordEntry()
	cForm.Append("Password", _password)
	cForm.Append("submit", widget.NewButton("create!Demo!Linode!", func() {
		displayFullScreenLog(displayLogNew("LinodeCreateDemo........."), window)
		err = LinodeApi.LinodeCreateDemo(_label.Text, _password.Text)
		displayFullScreenLog(displayLogAppend("LinodeCreateDemo finish:"+Common.SafeGetError(err)), window)
		log.Println("LinodeCreateDemo finish:" + Common.SafeGetError(err))
		log.Println(err)
		setTabContent(0, window)
	}))

	tab1 := widget.NewTabItemWithIcon("index", theme.HomeIcon(),
		widget.NewVBox(widget.NewLabel("Linode VPS Controller"),
			widget.NewHyperlink("github home page", u),
			cForm))
	tabs := widget.NewTabContainer(tab1)
	displayFullScreenLog(displayLogAppend("getting linode list"), window)
	infoList, err := getLinodeList()
	displayFullScreenLog(displayLogAppend("getting linode list finish:"+Common.SafeGetError(err)), window)

	tabID := 0
	if len(infoList) > 0 {
		for _, v := range infoList {
			tabID = tabID + 1
			tabs.Append(widget.NewTabItem(v.Label, createServerInfoForm(tabID, window, v)))
		}
	}

	log.Println(Common.FastJsonMarshal(infoList))

	widget.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab"))

	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(id)

	window.SetContent(tabs)
	//window.Resize(fyne.NewSize(600, 320))
	window.Show()
}

func displayFullScreenLog(logText string, window fyne.Window) {
	var logWidget = widget.NewLabel(logText)

	window.SetContent(widget.NewVBox(widget.NewProgressBarInfinite(), logWidget))

	window.Show()
}

func displayLogAppend(logText string) string {
	return textLog + "\n" + logText
}

func displayLogNew(logText string) string {
	textLog = logText
	return textLog
}

func getLinodeList() (infoList []LinodeApi.LinodeInfoJson, err error) {
	infoList, err = LinodeApi.LinodeList()
	return
}

func createServerInfoForm(tabID int, window fyne.Window, _info LinodeApi.LinodeInfoJson) *widget.Box {

	form := widget.NewForm()
	form.Append("DELETE", widget.NewButtonWithIcon("DELETE", theme.DeleteIcon(), func() {
		displayFullScreenLog(displayLogNew("deleting........."), window)
		LinodeApi.LinodeDelete(strconv.FormatInt(_info.ID, 10))
		displayFullScreenLog(displayLogAppend("deleting finish:"+Common.SafeGetError(err)), window)

		setTabContent(tabID, window)
		log.Println("delete..............................", err)
	}))

	form.Append("Created", widget.NewLabel(_info.Label))
	form.Append("status", widget.NewHBox(widget.NewLabel(_info.Status), widget.NewButtonWithIcon("Refresh", theme.ViewRefreshIcon(), func() {
		setTabContent(tabID, window)
		log.Println("refresh.................", tabID, err)

	})))

	form.Append("Created", widget.NewLabel(_info.Created))
	if "running" == _info.Status {
		form.Append("action", widget.NewButtonWithIcon("Stop", theme.CancelIcon(), func() {
			displayFullScreenLog(displayLogNew("Stoping........."), window)
			err = LinodeApi.LinodeShutdown(strconv.FormatInt(_info.ID, 10))
			displayFullScreenLog(displayLogAppend("Stoping finish:"+Common.SafeGetError(err)), window)
			setTabContent(tabID, window)

		}))
	} else {
		form.Append("action", widget.NewButtonWithIcon("Start", theme.CancelIcon(), func() {
			displayFullScreenLog(displayLogNew("Starting........."), window)
			err = LinodeApi.LinodeBoot(strconv.FormatInt(_info.ID, 10))
			displayFullScreenLog(displayLogAppend("Stoping finish:"+Common.SafeGetError(err)), window)
			setTabContent(tabID, window)
		}))
	}
	form.Append("Region", widget.NewLabel(_info.Region))
	form.Append("info", widget.NewLabel(_info.Image+Common.FastJsonMarshal(_info.Specs)))
	if len(_info.Ipv4) > 0 {
		for _, ip := range _info.Ipv4 {
			t := widget.NewEntry()
			t.Text = ip
			form.Append("ipv4", t)
		}
	}
	if len(_info.Ipv6) > 0 {
		t := widget.NewEntry()
		t.Text = _info.Ipv6
		form.Append("ipv6", t)
	}

	form.Append("action", widget.NewButtonWithIcon("Reboot", theme.ContentRedoIcon(), func() {
		LinodeApi.LinodeReBoot(strconv.FormatInt(_info.ID, 10))
		setTabContent(tabID, window)
		log.Println("reboot..............................", err)
	}))
	return widget.NewVBox(form)

}
