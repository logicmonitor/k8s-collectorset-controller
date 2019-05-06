package controller

import (
	"github.com/logicmonitor/lm-sdk-go/client"
)

func newLMClient(id, key, company string) *client.LMSdkGo {
	config := client.NewConfig()
	config.SetAccessID(&id)
	config.SetAccessKey(&key)
	domain := company + ".logicmonitor.com"
	config.SetAccountDomain(&domain)
	api := client.New(config)
	return api
}
