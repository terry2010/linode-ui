package LinodeApi

var token string
var tenantId string

//vps server ID list
var ServerIDList []string

var ApiURLs = struct {
	LinodeList string
	LinodeInfo string
	LinodeShutdown string
	LinodeBoot string
	LinodeReboot string
}{
	LinodeList: "https://api.linode.com/v4/linode/instances",
	LinodeInfo: "https://api.linode.com/v4/linode/instances/",
	LinodeShutdown: "https://api.linode.com/v4/linode/instances/%s/shutdown",
	LinodeBoot: "https://api.linode.com/v4/linode/instances/%s/boot",
	LinodeReboot: "https://api.linode.com/v4/linode/instances/%s/reboot",
}


var FileList = struct {
	Token string
}{
	Token: "token.json",
}

/**********************************************************************************************************************/
/************************************ linode api json struct **********************************************************/
/**********************************************************************************************************************/

type LinodeListJson struct {
	Data    []LinodeInfoJson `json:"data"`
	Page    int64            `json:"page"`
	Pages   int64            `json:"pages"`
	Results int64            `json:"results"`
}

type LinodeInfoJson struct {
	Alerts struct {
		CPU           int64 `json:"cpu"`
		Io            int64 `json:"io"`
		NetworkIn     int64 `json:"network_in"`
		NetworkOut    int64 `json:"network_out"`
		TransferQuota int64 `json:"transfer_quota"`
	} `json:"alerts"`
	Backups struct {
		Enabled        bool        `json:"enabled"`
		LastSuccessful interface{} `json:"last_successful"`
		Schedule       struct {
			Day    interface{} `json:"day"`
			Window interface{} `json:"window"`
		} `json:"schedule"`
	} `json:"backups"`
	Created    string   `json:"created"`
	Group      string   `json:"group"`
	Hypervisor string   `json:"hypervisor"`
	ID         int64    `json:"id"`
	Image      string   `json:"image"`
	Ipv4       []string `json:"ipv4"`
	Ipv6       string   `json:"ipv6"`
	Label      string   `json:"label"`
	Region     string   `json:"region"`
	Specs      struct {
		Disk     int64 `json:"disk"`
		Gpus     int64 `json:"gpus"`
		Memory   int64 `json:"memory"`
		Transfer int64 `json:"transfer"`
		Vcpus    int64 `json:"vcpus"`
	} `json:"specs"`
	Status          string        `json:"status"`
	Tags            []interface{} `json:"tags"`
	Type            string        `json:"type"`
	Updated         string        `json:"updated"`
	WatchdogEnabled bool          `json:"watchdog_enabled"`
}

type LinodeCreateJson struct {
	Booted   bool   `json:"booted"`
	Image    string `json:"image"`
	Label    string `json:"label"`
	Region   string `json:"region"`
	RootPass string `json:"root_pass"`
	SwapSize int64  `json:"swap_size"`
	Type     string `json:"type"`
	PrivateIP bool `json:"private_ip"`
}