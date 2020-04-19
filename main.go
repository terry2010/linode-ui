package main

import (
	"./common"
	"./linodeApi"
	"./gui"
)

var err error


func main() {
	err = Common.InitConfig()

	if nil != err {
		Gui.ErrorAlert(err.Error())
		return
	}


	err = LinodeApi.InitApi()
	if nil != err {
		Gui.ErrorAlert(err.Error())
		return
	}


	Gui.Run()


}
