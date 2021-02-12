package ethcloudflareclient

import (
	"net/http"
)

type EthCloudflareClient struct {
	HttpClient      *http.Client
	GetBlockByIdUrl string
}

func NewEthCloudflareClient(hc *http.Client, gu string) *EthCloudflareClient {
	return &EthCloudflareClient{
		HttpClient:      hc,
		GetBlockByIdUrl: gu,
	}
}
