package ethcloudflareclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/erizaver/etherium_proxy/internal/pkg/model"
)

const (
	GetBlockMethodName = "eth_getBlockByNumber"
)

//GetBlockByNumber will get block from cloudflare platform
func (c *EthCloudflareClient) GetBlockByNumber(blockID string) (*model.Block, error) {
	reqBody := GetBlockByNumberClientRequest{
		JsonRpc: "2.0",
		Method:  GetBlockMethodName,
		Params: []interface{}{
			blockID,
			true,
		},
		Id: 1,
	}
	marshalledReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return nil, errors.Wrap(err, "unable to marshal request")
	}

	httpRequest, err := http.NewRequest("POST", c.GetBlockByIdUrl, bytes.NewBuffer(marshalledReqBody))
	if err != nil {
		return nil, errors.Wrap(err, "unable to create http request")
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	resp, err := c.HttpClient.Do(httpRequest)
	if err != nil {
		return nil, errors.Wrap(err, "error making GetBlock request")
	}
	defer resp.Body.Close()

	httpResp := &GetBlockClientResponse{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read body")
	}
	fmt.Println(string(body))

	err = json.Unmarshal(body, httpResp)
	if err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal body")
	}

	return httpResp.Result, nil
}
