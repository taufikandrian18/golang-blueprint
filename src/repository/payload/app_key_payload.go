package payload

// ValidateAppKeyPayload ...
type ValidateAppKeyPayload struct {
	AppName string `json:"app_name"`
	AppKey  string `json:"app_key"`
}
