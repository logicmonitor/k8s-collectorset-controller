package lm_sdk_go_extension

/*func TestGetEscalationChainById(t *testing.T) {
	config := lm.NewConfiguration()
	config.APIKey = map[string]map[string]string{
		"Authorization": {
			"AccessID":  "***",
			"AccessKey": "***",
		},
	}
	config.BasePath = "https://qapr.logicmonitor.com/santaba/rest"

	api := NewExtensionApi(config)

	restResponse, apiResponse, err := api.GetEscalationChainById(31, "")
	if err != nil {
		t.Errorf("failed to get the escalation chain %v: %v", 31, err)
	}
	if apiResponse.StatusCode != http.StatusOK {
		t.Errorf("failed to get the escalation chain %v: %v", 31, apiResponse)
	}

	if &restResponse.Data == nil || restResponse.Data.Id != 31 {
		t.Errorf("failed to get the escalation chain %v: %v", 31, restResponse)
	}
}*/
