package LinodeApi

import (
	"../common"
	"fmt"
	"log"
)

func LinodeList() (infoList []LinodeInfoJson, err error) {
	apiURL := ApiURLs.LinodeList
	body, err := Common.HTTPGet(apiURL)
	fmt.Println(string(body))
	var llJson LinodeListJson

	Common.Json.Unmarshal(body, &llJson)
	infoList = llJson.Data
	return
}

func LinodeInfo(id string) {
	apiURL := ApiURLs.LinodeInfo + id
	body, err := Common.HTTPGet(apiURL)
	log.Println(err)
	fmt.Println(string(body))

}

func LinodeShutdown(id string) (err error) {
	apiURL := fmt.Sprintf(ApiURLs.LinodeShutdown, id)
	body, err := Common.HTTPPost(apiURL, "")
	log.Println(err)
	fmt.Println(string(body))
	return
}

func LinodeBoot(id string) (err error) {
	apiURL := fmt.Sprintf(ApiURLs.LinodeBoot, id)
	body, err := Common.HTTPPost(apiURL, "")
	log.Println(err)
	fmt.Println(string(body))
	return
}

func LinodeReBoot(id string) (err error) {
	apiURL := fmt.Sprintf(ApiURLs.LinodeReboot, id)
	body, err := Common.HTTPPost(apiURL, "")
	log.Println(err)
	fmt.Println(string(body))
	return
}

func LinodeCreateDemo(label string, password string) (err error) {
	apiURL := ApiURLs.LinodeList
	var data LinodeCreateJson
	data.Image = "linode/centos7"
	data.SwapSize = 512
	data.RootPass = password
	data.Booted = true
	data.Label = label
	data.Type = "g6-nanode-1"
	data.Region = "ap-northeast"
	data.PrivateIP = false
	log.Println("LinodeCreateDemoLinodeCreateDemoLinodeCreateDemoLinodeCreateDemoLinodeCreateDemoLinodeCreateDemoLinodeCreateDemoLinodeCreateDemoLinodeCreateDemoLinodeCreateDemoLinodeCreateDemo")
	log.Println(Common.FastJsonMarshal(data))
	body, err := Common.HTTPPost(apiURL, Common.FastJsonMarshal(data))
	log.Println(err)
	fmt.Println(string(body))
	return
}

func LinodeDelete(id string) (err error) {
	apiURL := ApiURLs.LinodeInfo + id
	body, err := Common.HTTPDelete(apiURL)
	log.Println(err)
	fmt.Println(string(body))
	return
}
