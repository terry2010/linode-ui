package LinodeApi

import (
	"../common"
	"github.com/pkg/errors"
	"time"
)

func InitApi() (err error) {
	Common.SetRetryTimes(Common.Config.GetInt("retrytime"))
	Common.SetTimeOut(time.Duration(Common.Config.GetInt("timeout")) * time.Second)
	token = Common.Config.GetString("token")
	if len(token) < 64 {
		return errors.New("token invalid")
	}
	Common.SetHeader("Authorization", "Bearer "+token)
	Common.SetHeader("Content-Type", "application/json")


	return

}
