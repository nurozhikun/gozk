package znet

import (
	"github.com/go-resty/resty/v2"
)

type HttpClient struct {
	*resty.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		Client: resty.New(),
	}
}

type (
	Response = resty.Response
	Request  = resty.Request
)
