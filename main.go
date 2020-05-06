package main

import (
	"./common"
	"./gui"
	"./linodeApi"
)

var err error

func main() {
	err = Common.InitConfig()

	if nil != err {
		Gui.InputToken(err.Error())
		return
	}

	err = LinodeApi.InitApi()

	if nil != err {
		Gui.InputToken(err.Error())
		return
	}

	Gui.Run()

}
