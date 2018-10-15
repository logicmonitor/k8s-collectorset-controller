package lm_sdk_go_extension

import (
	"encoding/json"
	"fmt"
	lm "github.com/logicmonitor/lm-sdk-go"
	"net/url"
	"strings"
)

type ExtensionApi struct {
	Configuration *lm.Configuration
}

func NewExtensionApi(config *lm.Configuration) *ExtensionApi {
	return &ExtensionApi{
		Configuration: config,
	}
}

/**
 * get collector
 *
 *
 * @param id
 * @param fields
 * @return *RestCollectorResponse
 */
func (a ExtensionApi) GetEscalationChainById(id int32, fields string) (*RestEscalationChainResponse, *lm.APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/setting/alert/chains/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", fmt.Sprintf("%v", id), -1)
	localResourcePath := "/setting/alert/chains/{id}"
	localResourcePath = strings.Replace("/setting/alert/chains/{id}", "{"+"id"+"}", fmt.Sprintf("%v", id), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("fields", a.Configuration.APIClient.ParameterToString(fields, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// authentication '(LMv1)' required
	// set key with prefix in header
	localVarHeaderParams["Authorization"] = a.Configuration.GetAPIKeyWithPrefix("Authorization", localResourcePath, "Get", localVarPostBody)
	var successPayload = new(RestEscalationChainResponse)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &lm.APIResponse{Operation: "GetEscalationChainById", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}
