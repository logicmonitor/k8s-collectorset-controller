package lm_sdk_go_extension

type RestEscalationChainResponse struct {
	Data RestEscalationChain `json:"data,omitempty"`

	Errmsg string `json:"errmsg,omitempty"`

	Status int32 `json:"status,omitempty"`
}
