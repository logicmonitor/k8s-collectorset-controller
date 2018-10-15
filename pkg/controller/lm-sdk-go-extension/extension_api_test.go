package lm_sdk_go_extension

import (
	lm "github.com/logicmonitor/lm-sdk-go"
	"net/http"
	"testing"
)

func TestGetEscalationChainById(t *testing.T) {
	config := lm.NewConfiguration()
	config.APIKey = map[string]map[string]string{
		"Authorization": {
			"AccessID":  "d8W84Vm67HzuR3YXX9BG",
			"AccessKey": "QDj^U%_y=mMSH5U9{KBa9N{DtkVzdt(^y5_4-T4R",
		},
	}
	config.BasePath = "https://qapr.logicmonitor.com/santaba/rest"

	api := NewExtensionApi(config)

	restResponse, apiResponse, err := api.GetEscalationChainById(31223, "")
	if err != nil {
		t.Errorf("failed to get the escalation chain %v: %v", 31, err)
	}
	if apiResponse.StatusCode != http.StatusOK {
		t.Errorf("failed to get the escalation chain %v: %v", 31, apiResponse)
	}

	if &restResponse.Data == nil || restResponse.Data.Id != 31 {
		t.Errorf("failed to get the escalation chain %v: %v", 31, restResponse)
	}
}
