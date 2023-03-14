package model

type InvokeInterface struct {
	AccessKey string `json:"access_key" binding:"required"`
	SecretKey string `json:"secret_key" binding:"required"`
	Body      string `json:"body"`
}

type InterfaceInfo struct {
	UserId         uint   `json:"user_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Url            string `json:"url"`
	Method         string `json:"method"`
	RequestHeader  string `json:"request_header"`
	ResponseHeader string `json:"response_header"`
	RequestParams  string `json:"request_params"`
	Status         bool   `json:"status"`
}
