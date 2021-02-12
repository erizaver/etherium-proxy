package ethcloudflareclient

import "github.com/erizaver/etherium_proxy/internal/pkg/model"

type GetBlockByNumberClientRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int           `json:"id"`
}

type GetBlockClientResponse struct {
	JsonRpc string      `json:"jsonrpc"`
	Id      int         `json:"id"`
	Result  model.Block `json:"result"`
}
