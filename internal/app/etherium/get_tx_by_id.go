package etherium

import (
	"context"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/erizaver/etherium_proxy/internal/pkg/model"
	"github.com/erizaver/etherium_proxy/pkg/api"
)

func (e *EthApi) GetTx(ctx context.Context, req *api.GetTxRequest) (*api.GetTxResponse, error) {
	if req.GetBlockId() == "" || req.GetTxId() == "" {
		return nil, errors.New("blockId or txId can`t be empty")
	}

	index, isIndex := isIndex(req.GetTxId())

	blockId := e.getSafeblockId(req.GetBlockId())
	if blockId == "" {
		return nil, errors.New("unable to parse block Id")
	}

	block, err := e.EthService.GetBlockByNumber(blockId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get block")
	}

	if isIndex {
		if len(block.Transactions) <= int(index) {
			return nil, errors.New("block does not contain this transaction index")
		}
		return &api.GetTxResponse{
			Transaction: model.CastModelTransactionToPb(block.Transactions[index]),
		}, nil
	} else {
		tx, ok := block.FastTransactions[req.GetTxId()]
		if ok {
			return &api.GetTxResponse{
				Transaction: model.CastModelTransactionToPb(tx),
			}, nil
		}
	}

	return nil, errors.New("unable to get this transaction from this block")
}

//isIndex will parse txId
func isIndex(txId string) (index int64, isIndex bool) {
	if id, err := strconv.ParseInt(txId, 10, 64); err == nil {
		return id, true
	}
	if id, err := strconv.ParseInt(strings.Replace(txId, "0x", "", -1), 16, 64); err == nil {
		return id, true
	}
	return 0, false
}
