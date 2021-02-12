package etherium

import (
	"context"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"github.com/pkg/errors"

	"github.com/erizaver/etherium_proxy/internal/pkg/model"
	"github.com/erizaver/etherium_proxy/pkg/api"
)

const latestBlockID = "latest"

//GetBlockByNumber will return proto structure with ethereum block by its id
func (e *EthApi) GetBlockByNumber(ctx context.Context, req *api.GetBlockByNumberRequest) (*api.GetBlockByNumberResponse, error) {
	if req.GetBlockId() == "" {
		return nil, errors.New("block ID can`t be empty")
	}

	block, err := e.EthService.GetBlockByNumber(e.getSafeBlockId(req.GetBlockId()))
	if err != nil {
		return nil, errors.Wrap(err, "unable to get block")
	}

	return &api.GetBlockByNumberResponse{
		Block: model.CastModelBlockToPb(block),
	}, nil
}

//getSafeBlockId will return block number in Hex(since cloudflare uses hex)
func (e *EthApi) getSafeBlockId(rawBlockID string) (safeBlockId string) {
	if strings.EqualFold(rawBlockID, latestBlockID) {
		return latestBlockID

	} else if strings.HasPrefix(rawBlockID, "0x") {
		_, err := strconv.ParseInt(strings.Replace(rawBlockID, "0x", "", -1), 16, 64)
		if err != nil {
			return ""
		}
		return rawBlockID

	} else if id, err := strconv.ParseInt(rawBlockID, 10, 64); err == nil {
		hexId := "0x" + strconv.FormatInt(id, 16)
		return hexId

	} else {
		return ""
	}
}

// Not sure if we need this, need more info about API usability, if latest block will be used less then in every 30 seconds, we dont need this
func (e *EthApi) WarmUpLatestBlockNumber() {
	ctx := context.Background()
	req := &api.GetBlockByNumberRequest{
		BlockId: latestBlockID,
	}

	//We can update every 30 seconds to get always the best last block counter and keep all info up-to-date
	//go func() {
	//	for {
	//		if _, err := e.GetBlockByNumber(ctx, req); err != nil {
	//			glog.Error("error while warming up last block number", err)
	//		}
	//		time.Sleep(30 * time.Second)
	//	}
	//}()

	//or we can update it once-per-app-run, just to get last block id once
	if _, err := e.GetBlockByNumber(ctx, req); err != nil {
		glog.Error("error while warming up last block number", err)
	}
}