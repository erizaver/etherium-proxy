package etherium

import (
	"github.com/erizaver/etherium_proxy/internal/pkg/model"
	"github.com/erizaver/etherium_proxy/pkg/api"
)

type EthFacade struct {
	EthService EthService
	api.UnimplementedEthServiceServer
}

func NewEthFacade(es EthService) *EthFacade {
	return &EthFacade{EthService: es}
}

type EthService interface {
	GetBlockByNumber(hexBlockId string) (model.Block, error)
}
