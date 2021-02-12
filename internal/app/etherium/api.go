package etherium

import (
	"github.com/erizaver/etherium_proxy/internal/pkg/model"
	"github.com/erizaver/etherium_proxy/pkg/api"
)

type EthApi struct {
	EthService EthService
	api.UnimplementedEthServiceServer
}

func NewEthApi(es EthService) *EthApi {
	return &EthApi{EthService: es}
}

type EthService interface {
	GetBlockByNumber(hexblockId string) (*model.Block, error)
}
