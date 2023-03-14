package model

type Header struct {
	AccessKey string
	Nonce     int
	Timestamp string
	Sign      string
	Body      string
}
